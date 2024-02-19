package interfaces

import (
	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type CategoryUseCase interface {
	AddCategory(category models.AddCategory) (domain.Category, error)
	EditCategory(EditCategory models.EditCategory, id int) (domain.Category, error)
	DeleteCategory(id int) error
	ListCategories() ([]domain.Category, error)
	FilterByCategory(categoryID int) ([]models.FilterByCategoryResponse, string, error)
}
