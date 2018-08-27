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
	Number      int    `json:"number" form:"number" query:"number"`
}

// Update function
// Update value of the Property
func (category *Category) Update() error {
	return db.Save(category).Error
}

// UpdateNotAll function
// Update value of the Property
func (category *Category) UpdateNotAll() error {
	return db.Model(category).Update(category).Error
}

// GetCategoryByID function
func GetCategoryByID(UUID string) (*Category, error) {
	categ := &Category{}
	err := db.Where("id = ?", UUID).First(categ).Error
	return categ, err
}

// Delete function
// Delete property
func (category *Category) Delete() error {
	err := DeleteByCategory(category.ID)
	if err != nil {
		return err
	}
	category, _ = GetCategory(category.ID)
	err = db.Delete(category).Error
	if err != nil {
		return err
	}
	categories, err := GetCategories()
	if err != nil {
		return err
	}
	for _, categoryR := range categories {
		if categoryR.Number > category.Number {
			categoryR.Number--
			categoryR.Update()
		}
	}
	return err
}

// Create function
// Add new property and add it in db
// Return new property
func (category *Category) Create(name, icon, description string) (*Category, error) {
	count, _ := GetCategoriesCount()
	category = &Category{
		Name:        name,
		Icon:        icon,
		ID:          generateUUID(),
		Description: description,
		Number:      count + 1,
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

// GetCategory function
// Return all categories
func GetCategory(id string) (*Category, error) {
	categories := &Category{}
	err := db.Where("id = ?", id).First(categories).Error
	return categories, err
}

// String function
func (category *Category) String() string {
	return fmt.Sprint(*category)
}

// GetCategoriesCount function
func GetCategoriesCount() (int, error) {
	count := 0
	err := db.Table("categories").Count(&count).Error
	return count, err
}
