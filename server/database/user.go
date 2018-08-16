package database

import "github.com/jinzhu/gorm"

// User structure
// Email - email of user
// Password - encrypted string
// Role - admin(1) or usual user(0)
type User struct {
	gorm.Model
	email    string `gorm:"type:text;not null"`
	password string `gorm:"type:text;not null"`
	role     int    `gorm:"type:integer;not null"`
}

// GetEmail function
// Return email of yout user
func (user *User) GetEmail() string {
	return user.email
}

// GetPassword function
// Return password of yout user
func (user *User) GetPassword() string {
	return user.password
}

// GetRole function
// Return role of yout user
func (user *User) GetRole() int {
	return user.role
}

// AddUser function
// Add new user and add it in db
// Return new user
func AddUser(email, password string, role int) *User {
	user := &User{
		email:    email,
		password: password,
		role:     role,
	}
	db.Create(user)
	return user
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
	var role int
	switch isAdmin {
	case true:
		role = 1
		break
	case false:
		role = 0
		break
	}
	user.role = int(role)
	db.Save(user)
}

// DeleteUser function
// Delete user
func (user *User) DeleteUser() {
	db.Delete(user)
}

// GetUsers function
// Return all users
func GetUsers() []*User {
	var users []*User
	db.Find(users)
	return users
}
