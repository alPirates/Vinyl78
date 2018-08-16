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

// GetKey function
// Return key of yout property
func (property *Property) GetKey() string {
	return property.key
}

// GetValue function
// Return value of yout property
func (property *Property) GetValue() string {
	return property.value
}

// SetProperty function
// Set value of the Property
func (property *Property) SetProperty(value string) {
	property.value = value
	db.Save(property)
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

// DeleteProperty function
// Delete property
func (property *Property) DeleteProperty() {
	db.Delete(property)
}

// GetProperties function
// Return all properties
func GetProperties() []*Property {
	var properties []*Property
	db.Find(properties)
	return properties
}
