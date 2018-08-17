package database

import "github.com/jinzhu/gorm"

// Property structure
// Key - name of this property
// Value - value of this property
type Property struct {
	gorm.Model
	Key   string `json:"key" form:"key" query:"key"`
	Value string `json:"value" form:"value" query:"value"`
}

// Update function
// Update value of the Property
func (property *Property) Update(value string) {
	db.Save(property)
}

// Delete function
// Delete property
func (property *Property) Delete() {
	db.Delete(property)
}

// Create function
// Add new property and add it in db
// Return new property
func (property *Property) Create(key, value string) *Property {
	property = &Property{
		Key:   key,
		Value: value,
	}
	db.Create(property)
	return property
}

// GetProperties function
// Return all properties
func GetProperties() []*Property {
	properties := []*Property{}
	db.Find(&properties)
	return properties
}
