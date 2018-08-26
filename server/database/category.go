package database

import "fmt"

// Category structure
// Name - name of this category
// Number - id of this category
// Icon - path to icon of this category
type Category struct {
	ID          string `json:"id" form:"id" query:"id"`
	Name        string `json:"name" form:"name" query:"name"`
	Icon        string `json:"icon" form:"icon" query:"icon"`
	Description string `json:"description" form:"description" query:"description"`
}

// Update function
// Update value of the Property
func (category *Category) Update() error {
	return db.Save(category).Error
}

// GetCategoryById function
func GetCategoryById(UUID string) (*Category, error) {
	categ := &Category{}
	err := db.Where("id = ?", UUID).First(categ).Error
	return categ, err
}

// UpdateNotAll function
// Update value of the Property
func (category *Category) UpdateNotAll() error {
	return db.Model(category).Update(category).Error
}

// Delete function
// Delete property
func (category *Category) Delete() error {
	err := DeleteByCategory(category.ID)
	if err != nil {
		return err
	}
	return db.Delete(category).Error
}

// Create function
// Add new property and add it in db
// Return new property
func (category *Category) Create(name, icon, description string) (*Category, error) {
	category = &Category{
		Name:        name,
		Icon:        icon,
		ID:          generateUUID(),
		Description: description,
	}
	err := db.Create(category).Error
	return category, err
}

// GetCategories function
// Return all categories
func GetCategories() ([]*Category, error) {
	categories := []*Category{}
	err := db.Find(&categories).Error
	return categories, err
}

// String function
func (category *Category) String() string {
	return fmt.Sprint(*category)
}
