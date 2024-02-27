package interfaces

import (
	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type CategoryRepository interface {
	AddCategory(category models.AddCategory) (domain.Category, error)
	CheckCategoryExist(categoryName string) (bool, error)
	CheckCategoryExistByID(id int) (bool, error)
	EditCategory(EditCategory models.EditCategory, id int) (domain.Category, error)
	GetCategoryByID(id int) (domain.Category, error)
	DeleteCategory(id int) error
	GetCategories(page, per_product int) ([]domain.Category, error)
	FilterByCategory(categoryID,page, per_product int) ([]models.FilterByCategoryResponse, string, error)
}
