package database

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Application structure
// Message - text to admin from userplication structure
// Name - the name of the user
// Phone - phone number of the user
// Email - email of the user
// Status - in process(0) - completed(1) - refused(2)
type Application struct {
	ID          string    `json:"id" form:"id" query:"id"`
	CreatedTime time.Time `json:"time" form:"time" query:"time"`
	Name        string    `json:"name" form:"name" query:"name"`
	Phone       string    `json:"phone" form:"phone" query:"phone"`
	Email       string    `json:"email" form:"email" query:"email"`
	Message     string    `json:"message" form:"message" query:"message"`
	Status      uint      `json:"status" form:"status" query:"status"`
}

// Update function
// Update all about application
func (application *Application) Update() error {
	return db.Save(application).Error
}

// Delete function
// Delete your application
func (application *Application) Delete() error {
	return db.Delete(application).Error
}

// Create function
// Add new application and add it in db
// Return new application
func (application *Application) Create(name, phone, email, message string) (*Application, error) {
	application = &Application{
		ID:          generateUUID(),
		Name:        name,
		Phone:       phone,
		Email:       email,
		Message:     message,
		Status:      0,
		CreatedTime: time.Now(),
	}
	return application, db.Create(application).Error
}

// GetApplications function
// Return all applications
func GetApplications(skip, limit uint) ([]*Application, error) {
	applications := []*Application{}
	err := db.Offset(skip).Limit(limit).Find(&applications).Error
	return applications, err
}

// GetAllApplications function
// Return all applications
func GetAllApplications() ([]*Application, error) {
	applications := []*Application{}
	err := db.Find(&applications).Error
	return applications, err
}

// String function
func (application *Application) String() string {
	return fmt.Sprint(*application)
}

// GetApplicationsCount function
func GetApplicationsCount() (int, error) {
	count := 0
	err := db.Table("applications").Count(&count).Error
	return count, err
}

// generateUUID function
func generateUUID() string {
	rand.Seed(time.Now().UTC().UnixNano())

	uuid := ""
	uuid += generatePartOfUUID(8) + "-"
	uuid += generatePartOfUUID(4) + "-"
	uuid += generatePartOfUUID(4) + "-"
	uuid += generatePartOfUUID(4) + "-"
	uuid += generatePartOfUUID(12)

	return uuid

}

func generatePartOfUUID(n int) string {
	uuidPart := ""
	for i := 0; i < n; i++ {
		randInt := rand.Intn(36)
		if randInt < 10 {
			uuidPart += strconv.Itoa(randInt)
		} else {
			uuidPart += string(rune(randInt + 87))
		}
	}
	return uuidPart
}
