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
	Status  uint   `json:"status" form:"status" query:"status"`
}

// Update function
// Update all about application
func (application *Application) Update() {
	db.Save(application)
}

// Delete function
// Delete your application
func (application *Application) Delete() {
	db.Delete(application)
}

// Create function
// Add new application and add it in db
// Return new application
func (application *Application) Create(name, phone, email, message string) *Application {
	application = &Application{
		Name:    name,
		Phone:   phone,
		Email:   email,
		Message: message,
		Status:  0,
	}
	db.Create(application)
	return application
}

// GetApplications function
// Return all applications
func GetApplications(skip, limit uint) []*Application {
	applications := []*Application{}
	db.Offset(skip).Limit(limit).Find(&applications)
	return applications
}

// GetAllApplications function
// Return all applications
func GetAllApplications() []*Application {
	applications := []*Application{}
	db.Find(&applications)
	return applications
}
