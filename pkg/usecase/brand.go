package usecase

import (
	"errors"

	"github.com/ahdaan98/pkg/domain"
	repo "github.com/ahdaan98/pkg/repository/interface"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type BrandUseCase struct {
	repo repo.BrandRepository
}

func NewBrandUseCase(repo repo.BrandRepository) interfaces.BrandUseCase {
	return &BrandUseCase{
		repo: repo,
	}
}

func (br *BrandUseCase) AddBrand(Brand models.AddBrand) (domain.Brand, error) {
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

	Exist, err := br.repo.CheckBrandExist(EditBrand.BrandName)
	if err != nil {
		return domain.Brand{}, err
	}

	if Exist {
		return domain.Brand{}, errors.New("brand already exist")
	}

	brand, err := br.repo.EditBrand(EditBrand, id)
	if err != nil {
		return domain.Brand{}, err
	}

	return brand, nil
}

func (br *BrandUseCase) DeleteBrand(id int) error {
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
