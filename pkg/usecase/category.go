package usecase

import (
	"errors"

	"github.com/ahdaan98/pkg/domain"
	repo "github.com/ahdaan98/pkg/repository/interface"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type CategoryUseCase struct {
	repo repo.CategoryRepository
}

func NewCategoryUseCase(repo repo.CategoryRepository) interfaces.CategoryUseCase {
	return &CategoryUseCase{
		repo: repo,
	}
}

func (cat *CategoryUseCase) AddCategory(category models.AddCategory) (domain.Category, error) {
	if category.CategoryName == ""{
		return domain.Category{},errors.New("category name cannot be empty")
	}

	Exist, err := cat.repo.CheckCategoryExist(category.CategoryName)
	if err != nil {
		return domain.Category{}, err
	}

	if Exist {
		return domain.Category{}, errors.New("category already exist")
	}

	addedCategory, err := cat.repo.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}

	return addedCategory, nil
}

func (cat *CategoryUseCase) EditCategory(EditCategory models.EditCategory, id int) (domain.Category, error) {
	if EditCategory.CategoryName == "" {
		return domain.Category{},errors.New("category name cannot be empty")
	}

	exist,err:=cat.repo.CheckCategoryExistByID(id)
	if err!=nil{
		return domain.Category{},err
	}

	if !exist{
		return domain.Category{},errors.New("category with this id not exist")
	}

	Exist, err := cat.repo.CheckCategoryExist(EditCategory.CategoryName)
	if err != nil {
		return domain.Category{}, err
	}

	if Exist {
		return domain.Category{}, errors.New("category already exist")
	}

	category, err := cat.repo.EditCategory(EditCategory, id)
	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (cat *CategoryUseCase) DeleteCategory(id int) error {
	Exist, err := cat.repo.CheckCategoryExistByID(id)
	if err != nil {
		return err
	}

	if !Exist {
		return errors.New("category is not exist with this id")
	}

	if err := cat.repo.DeleteCategory(id); err != nil {
		return err
	}

	return nil
}

func (cat *CategoryUseCase) ListCategories() ([]domain.Category, error) {
	Categories, err := cat.repo.GetCategories()
	if err != nil {
		return []domain.Category{}, err
	}

	return Categories, nil
}

func (i *CategoryUseCase) FilterByCategory(categoryID int) ([]models.FilterByCategoryResponse, string, error) {
	Exist, err := i.repo.CheckCategoryExistByID(categoryID)
	if err != nil {
		return []models.FilterByCategoryResponse{},"", err
	}

	if !Exist {
		return []models.FilterByCategoryResponse{},"", errors.New("category does not exist with this id")
	}

	cat, catName, err := i.repo.FilterByCategory(categoryID)

	if err != nil {
		return []models.FilterByCategoryResponse{},"", err
	}

	return cat, catName, nil
}
