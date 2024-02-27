package usecase

import (
	"errors"

	"github.com/ahdaan98/pkg/domain"
	helper "github.com/ahdaan98/pkg/helper/interfaces"
	repo "github.com/ahdaan98/pkg/repository/interface"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type CategoryUseCase struct {
	repo   repo.CategoryRepository
	helper helper.Helper
}

func NewCategoryUseCase(repo repo.CategoryRepository, helper helper.Helper) interfaces.CategoryUseCase {
	return &CategoryUseCase{
		repo:   repo,
		helper: helper,
	}
}

func (cat *CategoryUseCase) AddCategory(category models.AddCategory) (domain.Category, error) {
	if category.CategoryName == "" {
		return domain.Category{}, errors.New("category name cannot be empty")
	}

	tr := cat.helper.ContainOnlyLetters(category.CategoryName)

	if !tr {
		return domain.Category{}, errors.New("category name can only contain letters")
	}

	if len(category.CategoryName) < 4 {
		return domain.Category{}, errors.New("cateogory name should contain atleast 4 characters")
	}

	if len(category.CategoryName) > 6 {
		return domain.Category{}, errors.New("cateogory name can contain upto 6 characters only")
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
		return domain.Category{}, errors.New("category name cannot be empty")
	}

	exist, err := cat.repo.CheckCategoryExistByID(id)
	if err != nil {
		return domain.Category{}, err
	}

	
	if !exist {
		return domain.Category{}, errors.New("category with this id not exist")
	}

	tr := cat.helper.ContainOnlyLetters(EditCategory.CategoryName)

	if !tr {
		return domain.Category{}, errors.New("category name can only contain letters")
	}

	if len(EditCategory.CategoryName) < 4 {
		return domain.Category{}, errors.New("cateogory name should contain atleast 4 characters")
	}

	if len(EditCategory.CategoryName) > 6 {
		return domain.Category{}, errors.New("cateogory name can contain upto 6 characters only")
	}

	cate, err := cat.repo.GetCategoryByID(id)
	if err != nil {
		return domain.Category{}, err
	}


	if cate.CategoryName == EditCategory.CategoryName {
		return domain.Category{}, errors.New("category name is same with the previous one")
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
	if id <=0 {
		return errors.New("check value properly, id cannot be negative or zero")
	}
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

func (cat *CategoryUseCase) ListCategories(page, per_product int) ([]domain.Category, error) {
	Categories, err := cat.repo.GetCategories(page, per_product)
	if err != nil {
		return []domain.Category{}, err
	}

	return Categories, nil
}

func (i *CategoryUseCase) FilterByCategory(categoryID, page, per_product int) ([]models.FilterByCategoryResponse, string, error) {

	if categoryID < 1 || page < 1 || per_product < 1 {
		return []models.FilterByCategoryResponse{}, "", errors.New("check values properly, id cannot be negative")
	}

	if categoryID == 0 || page == 0 || per_product == 0 {
		return []models.FilterByCategoryResponse{}, "", errors.New("check values properly, id cannot be zero")
	}

	Exist, err := i.repo.CheckCategoryExistByID(categoryID)
	if err != nil {
		return []models.FilterByCategoryResponse{}, "", err
	}

	if !Exist {
		return []models.FilterByCategoryResponse{}, "", errors.New("category does not exist with this id")
	}

	cat, catName, err := i.repo.FilterByCategory(categoryID, page, per_product)

	if err != nil {
		return []models.FilterByCategoryResponse{}, "", err
	}

	return cat, catName, nil
}
