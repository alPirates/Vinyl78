package database

import (
	"fmt"
	"os"
)

// Image structure
// Nmae - filename
// Category - category of sticker
// Number - number in show
type Image struct {
	ID       string `json:"id" form:"id" query:"id"`
	Name     string `json:"name" form:"name" query:"name"`
	LinkedID string `json:"linked_id" form:"linked_id" query:"linked_id"`
	Number   uint   `json:"number" form:"number" query:"number"`
}

// Update function
// Update all about image
func (image *Image) Update() error {
	return db.Save(image).Error
}

// UpdateNotAll function
// Update all cahnges fields image
func (image *Image) UpdateNotAll() error {
	return db.Model(image).Update(image).Error
}

// DeleteBySticker function
// Delete image by linkedUUID
func (image *Image) DeleteBySticker(linkedID string) error {

	images, _ := GetImages(linkedID)

	for _, imageR := range images {
		imageR.DeleteFileImage()
	}

	return db.Where("linked_id = ?", linkedID).Delete(image).Error

}

// Delete function
// Delete image
func (image *Image) Delete() error {
	image, err := GetImage(image.ID)
	if err != nil {
		return err
	}
	err = db.Delete(image).Error
	if err != nil {
		return err
	}
	images, err := GetImages(image.LinkedID)
	if err != nil {
		return err
	}
	for _, imageR := range images {
		if imageR.Number > image.Number {
			imageR.Number--
			imageR.Update()
		}
		imageR.DeleteFileImage()
	}
	return err
}

// DeleteFileImage function
func (image *Image) DeleteFileImage() {
	os.Remove("../media/" + image.Name)
}

// GetImage function
func GetImage(id string) (*Image, error) {
	image := &Image{}
	err := db.Where("id = ?", id).First(image).Error
	return image, err
}

// Create function
// Create new image and add it in db
// Return new image
func (image *Image) Create(name string, linkedID string, number uint) (*Image, error) {
	image = &Image{
		Name:     name,
		LinkedID: linkedID,
		Number:   number,
		ID:       generateUUID(),
	}
	err := db.Create(image).Error
	if err != nil {
		return nil, err
	}
	return image, err
}

// GetAllImage function
// Return all images
func GetAllImage() ([]*Image, error) {
	images := []*Image{}
	err := db.Find(&images).Error
	return images, err
}

// GetImages function
// Return all images in category
func GetImages(linkedID string) ([]*Image, error) {
	images := []*Image{}
	err := db.Where("linked_id = ?", linkedID).Order("number").Find(&images).Error
	return images, err
}

// String function
func (image *Image) String() string {
	return fmt.Sprint(*image)
}

// GetImagesCount function
func GetImagesCount(linkedUUID string) (int, error) {
	count := 0
	err := db.Table("images").Where("linked_id = ?", linkedUUID).Count(&count).Error
	return count, err
}
