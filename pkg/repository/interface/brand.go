package interfaces

import (
	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type BrandRepository interface {
	AddBrand(Brand models.AddBrand) (domain.Brand, error)
	CheckBrandExist(brandName string) (bool, error)
	CheckBrandExistByID(id int) (bool, error)
	EditBrand(EditBrand models.EditBrand, id int) (domain.Brand, error)
	DeleteBrand(id int) error
	GetBrands() ([]domain.Brand, error)
	FilterByBrand(id,page, per_product int) ([]models.FilterByBrandResponse, string, error)
}
