package database

import (
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
	db.AutoMigrate(&User{}, &Property{}, &Application{})
	user := &User{}
	db.Where("role = ?", 1).First(user)
	if user.ID == 0 {
		user = user.Create("admin", "admin")
		user.Role = 1
		user.Update()
	}
	return err
}

// CloseConnection function
// Disconnect to postgres DB
func CloseConnection() {
	db.Close()
}
