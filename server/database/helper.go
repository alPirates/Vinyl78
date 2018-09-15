package database

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // we use postgres
)

// DB is Vinyl78 project DB
var (
	db *gorm.DB
)

// ExportDatabase function
// create local backup file in ./backups folder
func ExportDatabase() error {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	hr, min, sec := currentTime.Clock()
	dirPath := "./backups/"
	formatDate := fmt.Sprintf(
		"%v-%v-%v_%v-%v-%v",
		year, month, day,
		hr, min, sec,
	)

	a, err := exec.Command("pg_dump", "-Fc", "vinyl78", "-f", dirPath+formatDate+".gz").Output()
	fmt.Println("a is ", a)
	if err != nil {
		return err
	}
	return nil
}

// OpenConnection function
// Connect to postgres DB
// nameDB is name of your DB
// return error if db is not open
func OpenConnection(nameDB string) error {
	var err error

	if err != nil {
		fmt.Println("config erro")
		return err
	}

	// for production build use user=root and password='test' credentials
	db, err = gorm.Open(
		"postgres",
		"password='postgres' dbname="+nameDB+" sslmode=disable",
		// "user=root password='test' dbname="+nameDB+" sslmode=disable",
	)
	if err != nil {
		fmt.Println("db not opened")
		return err
	}

	db.AutoMigrate(&User{}, &Property{}, &Application{}, &Sticker{}, &Category{}, &Image{})

	// Create admin if not existed
	user := &User{}
	db.Where("role = ?", 1).First(user)
	if user.ID == "" {
		user, err = user.Create("admin@mail.ru", "admin")
		if err != nil {
			fmt.Println("admin not created")
		}
		user.Role = 1
		err = user.Update()
		if err != nil {
			fmt.Println("admin not created")
		}
	}

	// Create CaruselID if not existed
	property := &Property{}
	db.Where("key = ?", "carousel_id").First(property)
	if property.Key == "" {
		_, err = property.Create("carousel_id", generateUUID(), true)
		if err != nil {
			fmt.Println("property for carousel not created")
		}
	}

	return err
}

// CloseConnection function
// Disconnect to postgres DB
func CloseConnection() {
	db.Close()
}

// generateUUID function
func generateUUID() string {
	rand.Seed(time.Now().UTC().UnixNano())

	uuid := ""
	uuid += generatePartOfUUID(8) + "-"
	uuid += generatePartOfUUID(4) + "-"
	uuid += generatePartOfUUID(4) + "-"
	uuid += generatePartOfUUID(4) + "-"
	uuid += generatePartOfUUID(12)

	return uuid

}

func generatePartOfUUID(n int) string {
	uuidPart := ""
	for i := 0; i < n; i++ {
		randInt := rand.Intn(36)
		if randInt < 10 {
			uuidPart += strconv.Itoa(randInt)
		} else {
			uuidPart += string(rune(randInt + 87))
		}
	}
	return uuidPart
}

// CheckUUID function
func CheckUUID(id string) bool {
	if len(id) == 36 && strings.Count(id, "-") == 4 {
		return true
	}
	return false
}
