package database

import "fmt"

// Image structure
// Nmae - filename
// Category - category of sticker, -1 - carusel
// Number - number in show
type Image struct {
	ID     uint   `json:"ID" form:"ID" query:"ID" gson:"PRIMARY_KEY"`
	Name   string `json:"name" form:"name" query:"name"`
	Linked int    `json:"linked" form:"linked" query:"linked"`
	Number uint   `json:"number" form:"number" query:"number"`
}

// Update function
// Update all about image
func (image *Image) Update() error {
	return db.Save(image).Error
}

// Delete function
// Delete image
func (image *Image) Delete() error {
	err := db.Delete(image).Error
	if err != nil {
		return err
	}
	images, err := GetImages(image.Linked)
	if err != nil {
		return err
	}
	for i, image := range images {
		image.Number = uint(i + 1)
		err = image.Update()
		if err != nil {
			return err
		}
	}
	return err
}

// DeleteBySticker function
// Delete image by sticker id
func (image *Image) DeleteBySticker(linked int) error {
	return db.Where("linked = ?", linked).Delete(image).Error
}

// Create function
// Create new image and add it in db
// Return new image
func (image *Image) Create(name string, linked int, number uint) (*Image, error) {
	image = &Image{
		Name:   name,
		Linked: linked,
		Number: number,
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
func GetImages(linked int) ([]*Image, error) {
	images := []*Image{}
	err := db.Where("linked = ?", linked).Order("number").Find(&images).Error
	return images, err
}

// String function
func (image *Image) String() string {
	return fmt.Sprint(*image)
}

// GetImagesCount function
func GetImagesCount(linked int) (int, error) {
	count := 0
	err := db.Table("images").Where("linked = ?", linked).Count(&count).Error
	return count, err
}
