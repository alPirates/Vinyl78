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
	email    string `gorm:"type:text;not null"`
	password string `gorm:"type:text;not null"`
	role     uint   `gorm:"type:integer;not null"`
}

// GetID function
// Return ID of your user
func (user *User) GetID() uint {
	return user.Model.ID
}

// GetEmail function
// Return email of your user
func (user *User) GetEmail() string {
	return user.email
}

// GetPassword function
// Return password of your user
func (user *User) GetPassword() string {
	return user.password
}

// GetRole function
// Return role of your user
func (user *User) GetRole() uint {
	return user.role
}

// SetEmail function
// Set email in your user
func (user *User) SetEmail(email string) {
	user.email = email
	db.Save(user)
}

// SetPassword function
// Set password in your user
func (user *User) SetPassword(password string) {
	user.password = password
	db.Save(user)
}

// SetRole function
// Set role in your user
func (user *User) SetRole(isAdmin bool) {
	var role uint
	switch isAdmin {
	case true:
		role = 1
		break
	case false:
		role = 0
		break
	}
	user.role = role
	db.Save(user)
}

// Set function
// Set all about user
func (user *User) Set(email, password string, role uint) {
	user.email = email
	user.password = password
	user.role = role
	db.Save(user)
}

// Delete function
// Delete user
func (user *User) Delete() {
	db.Delete(user)
}

// AddUser function
// Add new user and add it in db
// Return new user
func AddUser(email, password string, isAdmin bool) *User {
	var role uint
	switch isAdmin {
	case true:
		role = 1
		break
	case false:
		role = 0
		break
	}
	user := &User{
		email:    email,
		password: password,
		role:     role,
	}
	db.Create(user)
	return GetUser(email, password)
}

// GetUser function
// Return user from email and password or nil
func GetUser(email, password string) *User {
	var user *User
	db.Where(&User{
		email:    email,
		password: password,
	}).First(user)
	return user
}

// GetUserFromEmail function
// Return user from email or nil
func GetUserFromEmail(email string) *User {
	var user *User
	db.Where(&User{
		email: email,
	}).First(user)
	return user
}

// GetUsers function
// Return all users
func GetUsers() []*User {
	var users []*User
	db.Find(users)
	return users
}
