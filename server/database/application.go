package database

import (
	"github.com/jinzhu/gorm"
)

// Application structure
// Name - the name of the user
// Phone - phone number of the user
// Email - email of the user
// Message - text to admin from user
// Status - in process(0) - completed(1) - refused(2)
type Application struct {
	gorm.Model
	Name    string `json:"name" form:"name" query:"name"`
	Phone   string `json:"phone" form:"phone" query:"phone"`
	Email   string `json:"email" form:"email" query:"email"`
	Message string `json:"message" form:"message" query:"message"`
	Status  uint
}

// Set function
// Set all about application
func (application *Application) Update() {
	db.Save(application)
}

// Delete function
// Delete your application
func (application *Application) Delete() {
	db.Delete(application)
}

// AddApplication function
// Add new application and add it in db
// Return new application
func AddApplication(name, phone, email, message string, status uint) *Application {
	application := &Application{
		Name:    name,
		Phone:   phone,
		Email:   email,
		Message: message,
		Status:  status,
	}
	db.Create(application)
	return application
}

// GetApplications function
// Return all applications
func GetApplications() []*Application {
	applications := []*Application{}
	db.Find(&applications)
	return applications
}
