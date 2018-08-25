package database

import "fmt"

// Property structure
// Key - name of this property
// Value - value of this property
type Property struct {
	ID         string `json:"id" form:"id" query:"id"`
	Key        string `json:"key" form:"key" query:"key"`
	Value      string `json:"value" form:"value" query:"value"`
	Permission int    `json:"permission" form:"permission" query:"permission"`
}

// Update function
// Update value of the Property
func (property *Property) Update() error {
	return db.Save(property).Error
}

// Delete function
// Delete property
func (property *Property) Delete() error {
	return db.Delete(property).Error
}

// Create function
// Add new property and add it in db
// Return new property
func (property *Property) Create(key, value string, permission int) (*Property, error) {
	property = &Property{
		Key:        key,
		Value:      value,
		ID:         generateUUID(),
		Permission: permission,
	}
	err := db.Create(property).Error
	return property, err
}

// GetProperties function
// Return all properties
func GetProperties() ([]*Property, error) {
	properties := []*Property{}
	err := db.Find(&properties).Error
	return properties, err
}

// GetPropertyPrivateAndPublic function
func GetPropertyPrivateAndPublic(key string) (*Property, error) {
	property := &Property{}
	err := db.Where("key = ?", key).First(&property).Error
	return property, err
}

// GetPropertyPublic function
func GetPropertyPublic(key string) (*Property, error) {
	property := &Property{}
	err := db.Where("key = ? AND perrmision = ?", key, 1).First(&property).Error
	return property, err
}

// String function
func (property *Property) String() string {
	return fmt.Sprint(*property)
}
