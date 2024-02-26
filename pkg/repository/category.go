package repository

import (
	"errors"

	"github.com/ahdaan98/pkg/domain"
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &CategoryRepository{
		DB: DB,
	}
}

func (cat *CategoryRepository) AddCategory(category models.AddCategory) (domain.Category, error) {
	var AddedCategory domain.Category

	query := `
	INSERT INTO categories (category_name)
	VALUES (?) RETURNING *
	`
	err := cat.DB.Raw(query, category.CategoryName).Scan(&AddedCategory).Error
	if err != nil {
		return domain.Category{}, err
	}

	return AddedCategory, nil
}

func (cat *CategoryRepository) CheckCategoryExist(categoryName string) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) FROM categories
	WHERE category_name = ?
	`

	err := cat.DB.Raw(query, categoryName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count >= 1 {
		return true, errors.New("category exist")
	}

	return false, nil
}

func (cat *CategoryRepository) CheckCategoryExistByID(id int) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) FROM categories
	WHERE id = ?
	`

	err := cat.DB.Raw(query, id).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count >= 1 {
		return true, nil
	}

	return false, nil
}

func (cat *CategoryRepository) EditCategory(EditCategory models.EditCategory, id int) (domain.Category, error) {
	var Updatedcategory domain.Category

	query := `
	UPDATE categories
	SET
	category_name = ?
	WHERE id = ?
	`

	if err := cat.DB.Exec(query, EditCategory.CategoryName, id).Error; err != nil {
		return domain.Category{}, err
	}

	if err := cat.DB.First(&Updatedcategory, id).Error; err != nil {
		return domain.Category{}, err
	}

	return Updatedcategory, nil
}

func (cat *CategoryRepository) DeleteCategory(id int) error {
	query := `
	DELETE FROM categories
	WHERE id = ?
	`

	if err := cat.DB.Exec(query, id).Error; err != nil {
		return err
	}

	return nil
}

func (cat *CategoryRepository) GetCategoryByID(id int) (domain.Category, error) {
	var category domain.Category

	query := `select * from categories where id = ?`

	if err := cat.DB.Raw(query, id).Scan(&category).Error; err != nil {
		return domain.Category{},err
	}

	return category,nil
}

func (cat *CategoryRepository) GetCategories() ([]domain.Category, error) {
	var categoryLists []domain.Category

	query := `
	SELECT * FROM categories
	`

	if err := cat.DB.Raw(query).Scan(&categoryLists).Error; err != nil {
		return []domain.Category{}, err
	}

	return categoryLists, nil
}

func (cat *CategoryRepository) FilterByCategory(categoryID int) ([]models.FilterByCategoryResponse, string, error) {

	var products []models.FilterByCategoryResponse

	query := `
	SELECT i.id AS product_id, i.product_name,i.brand_id, b.brand_name AS brand, i.stock, i.price
	FROM inventories i
	INNER JOIN categories c ON i.category_id = c.id
	INNER JOIN brands b ON i.brand_id = b.id
	WHERE c.id = ?
	`

	if err := cat.DB.Raw(query, categoryID).Scan(&products).Error; err != nil {
		return []models.FilterByCategoryResponse{}, "", err
	}

	var category string

	if err := cat.DB.Raw("SELECT category_name FROM categories WHERE id = ?", categoryID).Scan(&category).Error; err != nil {
		return []models.FilterByCategoryResponse{}, "", err
	}

	return products, category, nil
}
