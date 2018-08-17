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

// GetID function
// Return ID of your user
func (user *User) GetID() uint {
	return user.Model.ID
}

// GetEmail function
// Return email of your user
func (user *User) GetEmail() string {
	return user.Email
}

// GetPassword function
// Return password of your user
func (user *User) GetPassword() string {
	return user.Password
}

// GetRole function
// Return role of your user
func (user *User) GetRole() uint {
	return user.Role
}

// SetEmail function
// Set email in your user
func (user *User) SetEmail(email string) {
	user.Email = email
	db.Save(user)
}

// SetPassword function
// Set password in your user
func (user *User) SetPassword(password string) {
	user.Password = password
	db.Save(user)
}

// SetRole function
// Set role in your user
func (user *User) SetRole(isAdmin bool) {
	user.Role = boolToRole(isAdmin)
	db.Save(user)
}

// Set function
// Set all about user
func (user *User) Set(email, password string, isAdmin bool) {
	user.Email = email
	user.Password = password
	user.Role = boolToRole(isAdmin)
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
	user := &User{
		Email:    email,
		Password: password,
		Role:     boolToRole(isAdmin),
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

// CheckID function
// Return true if id existed
func CheckID(ID uint) bool {
	user := &User{}
	db.Where(&User{
		Model: gorm.Model{
			ID: ID,
		},
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

func boolToRole(isAdmin bool) uint {
	var role uint
	switch isAdmin {
	case true:
		role = 1
		break
	case false:
		role = 0
		break
	}
	return role
}
