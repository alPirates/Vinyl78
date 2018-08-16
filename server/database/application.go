package database

import (
	"time"

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
	name    string
	phone   string
	email   string
	message string
	status  uint
}

// GetID function
// Return ID of your application
func (application *Application) GetID() uint {
	return application.Model.ID
}

// GetCreationTime function
// Return time of your application's create
func (application *Application) GetCreationTime() time.Time {
	return application.Model.CreatedAt
}

// GetName function
// Return name of your application's user
func (application *Application) GetName() string {
	return application.name
}

// GetPhone function
// Return phone of your application's user
func (application *Application) GetPhone() string {
	return application.phone
}

// GetEmail function
// Return email of your application's user
func (application *Application) GetEmail() string {
	return application.email
}

// GetMessage function
// Return message of your application's user
func (application *Application) GetMessage() string {
	return application.message
}

// GetStatus function
// Return status of your application
func (application *Application) GetStatus() uint {
	return application.status
}

// SetName function
// Set name in your application's user
func (application *Application) SetName(name string) {
	application.name = name
	db.Save(application)
}

// SetPhone function
// Set phone in your application's user
func (application *Application) SetPhone(phone string) {
	application.phone = phone
	db.Save(application)
}

// SetEmail function
// Set email in your application's user
func (application *Application) SetEmail(email string) {
	application.email = email
	db.Save(application)
}

// SetMessage function
// Set message in your application's user
func (application *Application) SetMessage(message string) {
	application.message = message
	db.Save(application)
}

// SetStatus function
// Set status in your application's user
func (application *Application) SetStatus(status uint) {
	application.status = status
	db.Save(application)
}

// Set function
// Set all about application
func (application *Application) Set(name, phone, email, message string, status uint) {
	application.name = name
	application.phone = phone
	application.email = email
	application.message = message
	application.status = status
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
		name:    name,
		phone:   phone,
		email:   email,
		message: message,
		status:  status,
	}
	db.Create(application)
	return application
}

// GetApplications function
// Return all applications
func GetApplications() []*Application {
	var applications []*Application
	db.Find(applications)
	return applications
}
