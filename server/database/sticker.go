package database

import "fmt"

// Sticker structure
// Description - text
// Image - path for image
type Sticker struct {
	ID          uint   `json:"ID" form:"ID" query:"ID" gson:"PRIMARY_KEY"`
	Description string `json:"description" form:"description" query:"description"`
	Image       string `json:"image" form:"image" query:"image"`
	Category    uint   `json:"category" form:"category" query:"category"`
}

// Update function
// Update value of the Property
func (sticker *Sticker) Update() error {
	return db.Save(sticker).Error
}

// Delete function
// Delete property
func (sticker *Sticker) Delete() error {
	return db.Delete(sticker).Error
}

// Create function
// Add new sticker and add it in db
// Return new sticker
func (sticker *Sticker) Create(description, image string, category uint) (*Sticker, error) {
	sticker = &Sticker{
		Description: description,
		Image:       image,
		Category:    category,
	}
	err := db.Create(sticker).Error
	return sticker, err
}

// GetStickers function
// Return all stickers
// skip and limit
func GetStickers(skip, limit, category uint) ([]*Sticker, error) {
	sticker := []*Sticker{}
	err := db.Offset(skip).Limit(limit).Where("category = ?", category).Find(&sticker).Error
	return sticker, err
}

// GetAllStickers function
// Return all stickers
func GetAllStickers() ([]*Sticker, error) {
	sticker := []*Sticker{}
	err := db.Find(&sticker).Error
	return sticker, err
}

// GetStickersCount function
// get count of sticker in category
func GetStickersCount(category uint) (int, error) {
	count := 0
	err := db.Table("stickers").Where("category = ?", category).Count(&count).Error
	return count, err
}

// String function
func (sticker *Sticker) String() string {
	return fmt.Sprint(*sticker)
}
