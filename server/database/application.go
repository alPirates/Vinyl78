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
	Name    string `json:"name" form:"name" query:"name"`
	Phone   string `json:"phone" form:"phone" query:"phone"`
	Email   string `json:"email" form:"email" query:"email"`
	Message string `json:"message" form:"message" query:"message"`
	Status  uint
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
	return application.Name
}

// GetPhone function
// Return phone of your application's user
func (application *Application) GetPhone() string {
	return application.Phone
}

// GetEmail function
// Return email of your application's user
func (application *Application) GetEmail() string {
	return application.Email
}

// GetMessage function
// Return message of your application's user
func (application *Application) GetMessage() string {
	return application.Message
}

// GetStatus function
// Return status of your application
// Status - in process(0) - completed(1) - refused(2)
func (application *Application) GetStatus() uint {
	return application.Status
}

// SetName function
// Set name in your application's user
func (application *Application) SetName(name string) {
	application.Name = name
	db.Save(application)
}

// SetPhone function
// Set phone in your application's user
func (application *Application) SetPhone(phone string) {
	application.Phone = phone
	db.Save(application)
}

// SetEmail function
// Set email in your application's user
func (application *Application) SetEmail(email string) {
	application.Email = email
	db.Save(application)
}

// SetMessage function
// Set message in your application's user
func (application *Application) SetMessage(message string) {
	application.Message = message
	db.Save(application)
}

// SetStatus function
// Set status in your application's user
func (application *Application) SetStatus(status uint) {
	application.Status = status
	db.Save(application)
}

// Set function
// Set all about application
func (application *Application) Set(name, phone, email, message string, status uint) {
	application.Name = name
	application.Phone = phone
	application.Email = email
	application.Message = message
	application.Status = status
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
