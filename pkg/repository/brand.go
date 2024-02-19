package repository

import (
	"github.com/ahdaan98/pkg/domain"
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"gorm.io/gorm"
)

type BrandRepository struct {
	DB *gorm.DB
}

func NewBrandRepository(DB *gorm.DB) interfaces.BrandRepository {
	return &BrandRepository{
		DB: DB,
	}
}

func (br *BrandRepository) AddBrand(Brand models.AddBrand) (domain.Brand, error) {
	var AddedBrand domain.Brand

	query := `
	INSERT INTO brands (brand_name) 
	VALUES (?) RETURNING *
	`

	err := br.DB.Raw(query, Brand.BrandName).Scan(&AddedBrand).Error
	if err != nil {
		return domain.Brand{}, nil
	}

	return AddedBrand, nil
}

func (br *BrandRepository) CheckBrandExist(brandName string) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) FROM brands 
	WHERE brand_name = ?
	`

	err := br.DB.Raw(query, brandName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count >= 1 {
		return true, nil
	}

	return false, nil
}

func (br *BrandRepository) CheckBrandExistByID(id int) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) FROM brands 
	WHERE id = ?
	`

	err := br.DB.Raw(query, id).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count >= 1 {
		return true, nil
	}

	return false, nil
}

func (br *BrandRepository) EditBrand(EditBrand models.EditBrand, id int) (domain.Brand, error) {
	var UpdatedBrand domain.Brand

	query := `
	UPDATE brands
	SET
	brand_name = ?
	WHERE id = ?
	`

	if err := br.DB.Exec(query, EditBrand.BrandName, id).Error; err != nil {
		return domain.Brand{}, err
	}

	if err := br.DB.First(&UpdatedBrand, id).Error; err != nil {
		return domain.Brand{}, err
	}

	return UpdatedBrand, nil
}

func (br *BrandRepository) DeleteBrand(id int) error {
	query := `
	DELETE FROM brands
	WHERE id = ?
	`

	if err := br.DB.Exec(query, id).Error; err != nil {
		return err
	}

	return nil
}

func (br *BrandRepository) GetBrands() ([]domain.Brand, error) {
	var brandList []domain.Brand

	query := `
	SELECT * FROM brands
	`

	if err := br.DB.Raw(query).Scan(&brandList).Error; err != nil {
		return []domain.Brand{}, err
	}

	return brandList, nil
}

func (br *BrandRepository) FilterByBrand(id int) ([]models.FilterByBrandResponse, string, error) {

	var products []models.FilterByBrandResponse

	query := `
	SELECT i.id AS product_id, i.product_name, i.category_id, c.category_name AS category, i.stock, i.Price
	FROM inventories i
	INNER JOIN categories c ON i.category_id = c.id
	INNER JOIN brands b ON i.brand_id = b.id
	WHERE i.brand_id = ?
	`
	if err := br.DB.Raw(query, id).Scan(&products).Error; err != nil {
		return []models.FilterByBrandResponse{}, "", err
	}

	var brand string

	if err := br.DB.Raw("SELECT brand_name FROM brands WHERE id = ?", id).Scan(&brand).Error; err != nil {
		return []models.FilterByBrandResponse{}, "", err
	}

	return products, brand, nil

}
