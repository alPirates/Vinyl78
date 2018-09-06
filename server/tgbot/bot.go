package tgbot

import (
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/alPirates/Vinyl78/server/database"
	telega "gopkg.in/telegram-bot-api.v4"
)

var (
	bot    *telega.BotAPI
	toSend []*database.Application
)

// StartBot function
func StartBot() {
	var err error
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

func sendMessageByEmail(app *database.Application) {
	//@TODO read from config or something
	email := ""
	password := ""

	auth := smtp.PlainAuth("", email, password, "smtp.yandex.ru")
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
	err := smtp.SendMail("smtp.yandex.ru:25", auth, email, to, msg)
	if err != nil {
		log.Fatal(err)
	}

}
