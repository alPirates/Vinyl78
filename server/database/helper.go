package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // we use postgres
)

// DB is Vinyl78 project DB
var (
	db *gorm.DB
)

// OpenConnection function
// Connect to postgres DB
// nameDB is name of your DB
// return error if db is not open
func OpenConnection(nameDB string) error {

	var err error
	db, err = gorm.Open("postgres", "password='postgres' dbname="+nameDB+" sslmode=disable")
	if err != nil {
		fmt.Println("db not opened")
		return err
	}

	db.AutoMigrate(&User{}, &Property{}, &Application{}, &Sticker{}, &Category{})

	// Create admin if not existed
	user := &User{}
	db.Where("role = ?", 1).First(user)
	if user.ID == 0 {
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

	return err
}

// CloseConnection function
// Disconnect to postgres DB
func CloseConnection() {
	db.Close()
}
