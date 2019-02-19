package database

import (
	"fmt"
	"strconv"
)

// Property structure
// Key - name of this property
// Value - value of this property
// PRIVATE = FALSE
type Property struct {
	ID         string `json:"id" form:"id" query:"id"`
	Key        string `json:"key" form:"key" query:"key"`
	Value      string `json:"value" form:"value" query:"value"`
	Permission bool   `json:"permission" form:"permission" query:"permission"`
}

// Update function
// Update value of the Property
func (property *Property) Update() error {
	return db.Save(property).Error
}

// UpdateNotAll function
// Update value of the Property
func (property *Property) UpdateNotAll() error {
	return db.Model(property).Update(property).Error
}

// Delete function
// Delete property
func (property *Property) Delete() error {
	return db.Delete(property).Error
}

func (p *Property) Save() error {
	return db.Save(&p).Error
}

// DeleteKV function
// Delete property
func (property *Property) DeleteKV(key, value string) error {
	return db.Where("key = ? AND value = ?", key, value).Delete(property).Error
}

// Create function
// Add new property and add it in db
// Return new property
// private - false
func (property *Property) Create(key, value string, permission bool) (*Property, error) {
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

// GetPropertiesByKey function
func GetPropertiesByKey(key string) ([]*Property, error) {
	properties := []*Property{}
	err := db.Where("key = ?", key).Find(&properties).Error
	return properties, err
}

// CheckProperty function
// Return true if key and value existed
func CheckProperty(key, value string) bool {
	property := &Property{}
	db.Where(&Property{
		Key:   key,
		Value: value,
	}).First(property)
	if property.ID == "" {
		return false
	}
	return true
}

func IncrementSeoLink(key string) bool {
	property, err := GetPropertyPrivateAndPublic(key)
	if err != nil {
		return false
	}
	i, err := strconv.Atoi(property.Value)
	if err != nil {
		return false
	}
	i++
	property.Value = fmt.Sprintf("%d", i)
	fmt.Println("Test", property.Value)
	db.Save(property)
	return true
}

// GetPropertyPrivateAndPublic function
func GetPropertyPrivateAndPublic(key string) (*Property, error) {
	property := &Property{}
	err := db.Where("key = ?", key).First(property).Error
	return property, err
}

// GetPropertyPublic function
func GetPropertyPublic(key string) (*Property, error) {
	property := &Property{}
	err := db.Where("key = ? AND permission = ?", key, true).First(property).Error
	return property, err
}

// String function
func (property *Property) String() string {
	return fmt.Sprint(*property)
}
