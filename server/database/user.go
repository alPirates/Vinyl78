package database

import (
	"github.com/jinzhu/gorm"
)

// User structure
// Email - email of user
// Password - encrypted string
// Role - admin(1) or usual user(0)
type User struct {
	gorm.Model
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Role     uint
}

// Update function
// Update all about user
func (user *User) Update() {
	db.Save(user)
}

// Delete function
// Delete user
func (user *User) Delete() {
	db.Delete(user)
}

// Create function
// Create new user and add it in db
// Return new user
func (user *User) Create(email, password string) *User {
	user = &User{
		Email:    email,
		Password: password,
		Role:     0,
	}
	db.Create(user)
	return GetUser(email, password)
}

// GetUser function
// Return user from email and password or nil
func GetUser(email, password string) *User {
	user := &User{}
	db.Where(&User{
		Email:    email,
		Password: password,
	}).First(user)
	return user
}

// GetUserByID function
// Return user from ID
func GetUserByID(ID uint) *User {
	user := &User{}
	db.First(user, ID)
	return user
}

// CheckEmailAndPassword function
// Return true if email and Password existed
func CheckEmailAndPassword(email, password string) bool {
	user := &User{}
	db.Where(&User{
		Email:    email,
		Password: password,
	}).First(user)
	if user.Email == "" {
		return false
	}
	return true
}

// GetUsers function
// Return all users
func GetUsers() []*User {
	users := []*User{}
	db.Find(&users)
	return users
}
