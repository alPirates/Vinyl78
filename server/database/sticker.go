package database

import "fmt"

// Sticker structure
// Description - text
// Image - path for image
type Sticker struct {
	ID          string   `json:"id" form:"id" query:"id"`
	Description string   `json:"description" form:"description" query:"description"`
	Images      []*Image `json:"images" form:"images" query:"images" gorm:"-"`
	CategoryID  string   `json:"category_id" form:"category_id" query:"category_id"`
}

// Update function
// Update value of the Property
func (sticker *Sticker) Update() error {
	return db.Save(sticker).Error
}

// Delete function
// Delete property
func (sticker *Sticker) Delete() error {
	err := (&Image{}).DeleteBySticker(sticker.ID)
	if err != nil {
		return err
	}
	return db.Delete(sticker).Error
}

// Create function
// Add new sticker and add it in db
// Return new sticker
func (sticker *Sticker) Create(description, categoryID string) (*Sticker, error) {
	sticker = &Sticker{
		Description: description,
		CategoryID:  categoryID,
		ID:          generateUUID(),
	}
	err := db.Create(sticker).Error
	return sticker, err
}

// GetStickers function
// Return all stickers
// skip and limit
func GetStickers(skip, limit uint, categoryUUID string) ([]*Sticker, error) {
	sticker := []*Sticker{}
	err := db.Offset(skip).Limit(limit).Where("category_id = ?", categoryUUID).Find(&sticker).Error
	return sticker, err
}

// GetAllStickersCategory function
// Return all stickers by category
func GetAllStickersCategory(categoryUUID string) ([]*Sticker, error) {
	sticker := []*Sticker{}
	err := db.Where("category_id = ?", categoryUUID).Find(&sticker).Error
	return sticker, err
}

// DeleteByCategory function
// Delete all stickers in this category
func DeleteByCategory(categoryUUID string) error {
	stickers, err := GetAllStickersCategory(categoryUUID)
	if err != nil {
		return err
	}
	for _, sticker := range stickers {
		err = (&Image{}).DeleteBySticker(sticker.ID)
		if err != nil {
			return err
		}
	}
	err = db.Delete(Sticker{}, "category_id = ?", categoryUUID).Error
	return err
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
func GetStickersCount(categoryUUID string) (int, error) {
	count := 0
	err := db.Table("stickers").Where("category_id = ?", categoryUUID).Count(&count).Error
	return count, err
}

// String function
func (sticker *Sticker) String() string {
	return fmt.Sprint(*sticker)
}

// GetPreStickers fucntion
// return limit for each category
func GetPreStickers(limit uint) ([]*Sticker, error) {
	stickers := []*Sticker{}
	err := db.Table("stickers").Limit(limit).Group("category_id").Find(&stickers).Error
	return stickers, err
}
