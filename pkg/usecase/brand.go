package usecase

import (
	"errors"

	"github.com/ahdaan98/pkg/domain"
	helper "github.com/ahdaan98/pkg/helper/interfaces"
	repo "github.com/ahdaan98/pkg/repository/interface"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type BrandUseCase struct {
	repo repo.BrandRepository
	helper helper.Helper
}

func NewBrandUseCase(repo repo.BrandRepository, helper helper.Helper) interfaces.BrandUseCase {
	return &BrandUseCase{
		repo: repo,
		helper: helper,
	}
}

func (br *BrandUseCase) AddBrand(Brand models.AddBrand) (domain.Brand, error) {
	if Brand.BrandName == "" {
		return domain.Brand{},errors.New("brand name cannot be empty")
	}

	tr := br.helper.ContainOnlyLetters(Brand.BrandName)

	if !tr {
		return domain.Brand{}, errors.New("brand name can only contain letters")
	}

	if len(Brand.BrandName) < 4 {
		return domain.Brand{}, errors.New("brand name should contain atleast 4 characters")
	}

	if len(Brand.BrandName) > 6 {
		return domain.Brand{}, errors.New("brand name can contain upto 6 characters only")
	}


	Exist, err := br.repo.CheckBrandExist(Brand.BrandName)
	if err != nil {
		return domain.Brand{}, err
	}

	if Exist {
		return domain.Brand{}, errors.New("brand already exist")
	}

	addedBrand, err := br.repo.AddBrand(Brand)
	if err != nil {
		return domain.Brand{}, err
	}

	return addedBrand, nil
}

func (br *BrandUseCase) EditBrand(EditBrand models.EditBrand, id int) (domain.Brand, error) {
	if EditBrand.BrandName == "" {
		return domain.Brand{},errors.New("brand name cannot be empty")
	}


	Exist, err := br.repo.CheckBrandExist(EditBrand.BrandName)
	if err != nil {
		return domain.Brand{}, err
	}

	if Exist {
		return domain.Brand{}, errors.New("brand already exist")
	}

	tr := br.helper.ContainOnlyLetters(EditBrand.BrandName)

	if !tr {
		return domain.Brand{}, errors.New("brand name can only contain letters")
	}

	if len(EditBrand.BrandName) < 4 {
		return domain.Brand{}, errors.New("brand name should contain atleast 4 characters")
	}

	if len(EditBrand.BrandName) > 6 {
		return domain.Brand{}, errors.New("brand name can contain upto 6 characters only")
	}


	brand, err := br.repo.EditBrand(EditBrand, id)
	if err != nil {
		return domain.Brand{}, err
	}

	return brand, nil
}

func (br *BrandUseCase) DeleteBrand(id int) error {
	exist,err:=br.repo.CheckBrandExistByID(id)
	if err!=nil{
		return err
	}

	if !exist{
		return errors.New("brand with this id does not exist")
	}
	if err := br.repo.DeleteBrand(id); err != nil {
		return err
	}

	return nil
}

func (br *BrandUseCase) ListBrands() ([]domain.Brand, error) {
	brands, err := br.repo.GetBrands()
	if err != nil {
		return []domain.Brand{}, err
	}

	return brands, nil
}

func (br *BrandUseCase) FilterByBrand(BrandID int) ([]models.FilterByBrandResponse, string, error) {
	Exist, err := br.repo.CheckBrandExistByID(BrandID)
	if err != nil {
		return []models.FilterByBrandResponse{}, "", err
	}

	if !Exist {
		return []models.FilterByBrandResponse{}, "", errors.New("brand does not exist with this id")
	}

	products, brandName, err := br.repo.FilterByBrand(BrandID)
	if err != nil {
		return []models.FilterByBrandResponse{}, "", err
	}

	return products, brandName, nil
}
