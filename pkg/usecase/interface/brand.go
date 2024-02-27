package interfaces

import (
	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type BrandUseCase interface {
	AddBrand(Brand models.AddBrand) (domain.Brand, error)
	EditBrand(EditBrand models.EditBrand, id int) (domain.Brand, error)
	DeleteBrand(id int) error
	ListBrands() ([]domain.Brand, error)
	FilterByBrand(BrandID,page,per_product int) ([]models.FilterByBrandResponse, string, error)
}