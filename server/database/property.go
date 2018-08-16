package database

import "github.com/jinzhu/gorm"

// Property structure
// Key - name of this property
// Value - value of this property
type Property struct {
	gorm.Model
	key   string `gorm:"type:text;not null"`
	value string `gorm:"type:text;not null"`
}

// GetID function
// Return ID of your property
func (property *Property) GetID() uint {
	return property.Model.ID
}

// GetKey function
// Return key of your property
func (property *Property) GetKey() string {
	return property.key
}

// GetValue function
// Return value of your property
func (property *Property) GetValue() string {
	return property.value
}

// SetValue function
// Set value of the Property
func (property *Property) SetValue(value string) {
	property.value = value
	db.Save(property)
}

// Delete function
// Delete property
func (property *Property) Delete() {
	db.Delete(property)
}

// AddProperty function
// Add new property and add it in db
// Return new property
func AddProperty(key, value string) *Property {
	property := &Property{
		key:   key,
		value: value,
	}
	db.Create(property)
	return property
}

// GetProperties function
// Return all properties
func GetProperties() []*Property {
	var properties []*Property
	db.Find(properties)
	return properties
}
