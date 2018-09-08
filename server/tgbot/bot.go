package tgbot

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"

	"github.com/alPirates/Vinyl78/server/config"
	"github.com/alPirates/Vinyl78/server/database"
	telega "gopkg.in/telegram-bot-api.v4"
)

var (
	bot    *telega.BotAPI
	toSend []*database.Application
	conf   config.Config
)

// StartBot function
func StartBot() {
	var err error

	err = conf.ReadConfig()
	if err != nil {
		fmt.Println("can't read config ", err)
		return
	}

	bot, err = telega.NewBotAPI("689989185:AAELI2tILkC6d_feyy-sXpBFvHg-5oFUqe8")
	if err != nil {
		fmt.Println("can't start bot ", err)
		return
	}

	fmt.Println("")
	fmt.Println(
		`  __   ___  ___  ___  ___  __   ___  ___`+"\n",
		`|  | /  / /  / /   |/  / |  |_/  / /  /`+"\n",
		`|  |/  / /  / /       /  |_    _/ /  /__`+"\n",
		`|_____/ /__/ /__/|___/    /___/  /_____/`,
	)
	fmt.Println("Bot for web site Vinyl78 @Vinyl78Bot")

	workUpdate()
}

func workUpdate() {
	u := telega.NewUpdate(0)
	u.Timeout = 30

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Command() != "" {
			switch update.Message.Command() {
			case "start":
				bot.Send(telega.NewMessage(
					update.Message.Chat.ID,
					"/help - все команды\n/auth <email> <password> - авторизация\n/unauth - выход",
				))
				break
			case "help":
				bot.Send(telega.NewMessage(
					update.Message.Chat.ID,
					"/help - все команды\n/auth <email> <password> - авторизация\n/unauth - выход",
				))
				break
			case "unauth":
				(&database.Property{}).DeleteKV(
					"chat_id_telegram",
					fmt.Sprint(update.Message.Chat.ID),
				)
				bot.Send(telega.NewMessage(
					update.Message.Chat.ID,
					"вы вышли",
				))
				break
			case "auth":
				mas := strings.Split(update.Message.CommandArguments(), " ")
				if len(mas) == 2 {
					email := mas[0]
					password := mas[1]
					user, err := database.GetUser(email, password)
					if err != nil || user.Role == 0 {
						bot.Send(telega.NewMessage(
							update.Message.Chat.ID,
							"вы не админ",
						))
						break
					}
					if !database.CheckProperty("chat_id_telegram", fmt.Sprint(update.Message.Chat.ID)) {
						(&database.Property{}).Create(
							"chat_id_telegram",
							fmt.Sprint(update.Message.Chat.ID),
							false,
						)
					}
					bot.Send(telega.NewMessage(
						update.Message.Chat.ID,
						"теперь все новые заявки будут приходить сюда",
					))
				} else {
					bot.Send(telega.NewMessage(
						update.Message.Chat.ID,
						"неправильные аргументы функции",
					))
					break
				}
				break
			default:
				bot.Send(telega.NewMessage(
					update.Message.Chat.ID,
					"нет такой команды",
				))
				break
			}
		}
	}
}

// SendApplication function
func SendApplication(app *database.Application) {
	properties, err := database.GetPropertiesByKey("chat_id_telegram")
	if err != nil {
		sendMessageByEmail(app)
		return
	}
	for _, property := range properties {
		id, _ := strconv.ParseInt(property.Value, 10, 64)
		_, err = bot.Send(telega.NewMessage(
			id,
			app.Name+"\n"+app.Email+"\n"+app.Phone+"\n"+app.Message,
		))
		if err != nil {
			sendMessageByEmail(app)
			return
		}
	}
}

func getSmtpName(name string) (map[string]string, error) {
	servers := map[string]map[string]string{
		"yandex": map[string]string{
			"smtp": "smtp.yandex.ru",
			"port": "smtp.yandex.ru:25",
		},
		"google": map[string]string{
			"smtp": "smtp.gmail.com",
			"port": "smtp.gmail.com:587",
		},
	}
	rgx := regexp.MustCompile(`@(\w+)`)
	output := rgx.FindStringSubmatch(name)
	if len(output) > 1 {
		r := servers[output[1]]
		if r != nil {
			return r, nil
		}
		return nil, errors.New("Cant find smtp server in list")
	}
	return nil, errors.New("Cant find email addres regexp in string")
}

func sendMessageByEmail(app *database.Application) {
	//@TODO red from config or something
	email := conf.Email
	password := conf.Password
	serverSMTP, err := getSmtpName(email)

	if err != nil {
		log.Println(err)
		return
	}

	auth := smtp.PlainAuth("", email, password, serverSMTP["smtp"])
	to := []string{email}
	msg := []byte(
		"Subject: Новая заявка!\r\n" +
			"\r\n" +
			fmt.Sprintf(
				"Телефон: %s\nИмя: %s\nE-mail: %s\nСообщение: %s\n",
				app.Phone,
				app.Name,
				app.Email,
				app.Message,
			) +
			"\r\n",
	)
	err = smtp.SendMail(serverSMTP["port"], auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
	}

}
