package database

import "fmt"

// User structure
// Email - email of user
// Password - encrypted string
// Role - admin(1) or usual user(0)
type User struct {
	ID       uint   `json:"ID" form:"ID" query:"ID" gson:"PRIMARY_KEY"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Role     uint
}

// Update function
// Update all about user
func (user *User) Update() error {
	return db.Save(user).Error
}

// Delete function
// Delete user
func (user *User) Delete() error {
	return db.Delete(user).Error
}

// Create function
// Create new user and add it in db
// Return new user
func (user *User) Create(email, password string) (*User, error) {
	user = &User{
		Email:    email,
		Password: password,
		Role:     0,
	}
	err := db.Create(user).Error
	if err != nil {
		return nil, err
	}
	user, err = GetUser(email, password)
	return user, err
}

// GetUser function
// Return user from email and password or nil
func GetUser(email, password string) (*User, error) {
	user := &User{}
	err := db.Where(&User{
		Email:    email,
		Password: password,
	}).First(user).Error
	return user, err
}

// GetUserByID function
// Return user from ID
func GetUserByID(ID uint) (*User, error) {
	user := &User{}
	err := db.First(user, ID).Error
	return user, err
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
func GetUsers() ([]*User, error) {
	users := []*User{}
	err := db.Find(&users).Error
	return users, err
}

// String function
func (user *User) String() string {
	return fmt.Sprint(*user)
}
