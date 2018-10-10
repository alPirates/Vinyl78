package database

import (
	"fmt"
)

// Carousel structure
// Nmae - filename
// Category - category of sticker
// Number - number in show
type Carousel struct {
	ID     string `json:"id" form:"id" query:"id"`
	Name   string `json:"name" form:"name" query:"name"`
	Image  *Image `json:"image" form:"image" query:"-"`
	Number uint   `json:"number" form:"number" query:"number"`
}

func (carousel *Carousel) Update() error {
	return db.Save(carousel).Error
}

// UpdateNotAll function
// Update value of the Property
func (carousel *Carousel) UpdateNotAll() error {
	return db.Model(carousel).Update(carousel).Error
}

// Delete function
// Delete your application
func (carousel *Carousel) Delete() error {
	carousel.Image.Delete()
	return db.Delete(carousel).Error
}

// Create function
// Add new application and add it in db
// Return new application
func (carousel *Carousel) Create(name string, number uint) (*Carousel, error) {
	carousel = &Carousel{
		ID:     generateUUID(),
		Name:   name,
		Number: number,
	}
	return carousel, db.Create(carousel).Error
}

// GetApplications function
// Return all applications
func GetCarousel(skip, limit uint) ([]*Carousel, error) {
	carousels := []*Carousel{}
	err := db.Offset(skip).Limit(limit).Find(&carousels).Error
	return carousels, err
}

// GetAllApplications function
// Return all applications
func GetAllCarousel() ([]*Carousel, error) {
	carousels := []*Carousel{}
	err := db.Find(&carousels).Error
	return carousels, err
}

// String function
func (carousel *Carousel) String() string {
	return fmt.Sprint(*carousel)
}

// GetApplicationsCount function
func GetCarouselCount() (int, error) {
	count := 0
	err := db.Table("carousel").Count(&count).Error
	return count, err
}
