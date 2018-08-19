package database

import "fmt"

// Property structure
// Key - name of this property
// Value - value of this property
type Property struct {
	ID    uint   `json:"ID" form:"ID" query:"ID" gson:"PRIMARY_KEY"`
	Key   string `json:"key" form:"key" query:"key"`
	Value string `json:"value" form:"value" query:"value"`
}

// Update function
// Update value of the Property
func (property *Property) Update() error {
	return db.Model(property).Where("key = ?", property.Key).Update("value", property.Value).Error
}

// Delete function
// Delete property
func (property *Property) Delete() error {
	return db.Delete(property).Error
}

// Create function
// Add new property and add it in db
// Return new property
func (property *Property) Create(key, value string) (*Property, error) {
	property = &Property{
		Key:   key,
		Value: value,
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

// String function
func (property *Property) String() string {
	return fmt.Sprint(*property)
}
