package repository

import (
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"gorm.io/gorm"
)

type InventoryRepostiory struct {
	DB *gorm.DB
}

func NewInventoryRespository(DB *gorm.DB) interfaces.InventoryRepository {
	return &InventoryRepostiory{
		DB: DB,
	}
}

func (inv *InventoryRepostiory) UploadImage(id int, image string) error {
	if err := inv.DB.Exec("INSERT INTO images (inventory_id,image) VALUES (?,?)", id, image).Error; err != nil {
		return err
	}

	return nil
}

func (inv *InventoryRepostiory) AddInventory(inventory models.AddInventory) (models.InventoryResponse, error) {
	var ReturningInventories models.InventoryResponse
	query := `
	INSERT INTO inventories (product_name,brand_id,category_id,stock,price)
	VALUES (?,?,?,?,?)
	`
	err := inv.DB.Exec(query, inventory.ProductName, inventory.BrandID, inventory.CategoryID, inventory.Stock, inventory.Price).Error
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return ReturningInventories, nil
}

func (inv *InventoryRepostiory) CheckInventoryExist(productName string) (bool, error) {
	var count int
	query := `
	SELECT COUNT(*) FROM inventories
	WHERE product_name = ?
	`
	err := inv.DB.Raw(query, productName).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count >= 1 {
		return true, nil
	}

	return false, nil
}

func (inv *InventoryRepostiory) CheckInventoryExistByID(id int) (bool, error) {
	var count int

	query := `
	SELECT COUNT(*) FROM inventories
	WHERE id = ?
	`

	err := inv.DB.Raw(query, id).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count >= 1 {
		return true, nil
	}

	return false, nil
}

func (inv *InventoryRepostiory) ListProducts(page, per_product int) ([]models.InventoryResponse, error) {
	var productLists []models.InventoryResponse

	query := `
	SELECT i.id AS product_id, i.product_name, i.category_id, c.category_name AS category, i.brand_id, b.brand_name AS brand, i.stock, i.price
   FROM inventories i
   INNER JOIN categories c ON i.category_id = c.id
   INNER JOIN brands b ON i.brand_id = b.id
  `

	err := inv.DB.Raw(query).Scan(&productLists).Error
	if err != nil {
		return []models.InventoryResponse{}, err
	}

	return productLists, nil

}

func (inv *InventoryRepostiory) EditInventory(inventory models.EditInventory, id int) (models.InventoryResponse, error) {

	query := `
	UPDATE inventories
	SET product_name = ?, brand_id = ?, category_id = ?, price = ?
	WHERE id = ?
	`

	if err := inv.DB.Exec(query, inventory.ProductName, inventory.BrandID, inventory.CategoryID, inventory.Price, id).Error; err != nil {
		return models.InventoryResponse{}, err
	}

	var inventoryResponse models.InventoryResponse

	if err := inv.DB.Raw(`SELECT i.id AS product_id, i.product_name, i.category_id, c.category_name AS category, i.brand_id, b.brand_name AS brand, i.stock, i.price
  FROM inventories i
  INNER JOIN categories c ON i.category_id = c.id
  INNER JOIN brands b ON i.brand_id = b.id WHERE i.id = ?`, id).Scan(&inventoryResponse).Error; err != nil {
		return models.InventoryResponse{}, err
	}

	return inventoryResponse, nil
}

func (inv *InventoryRepostiory) UpdateInventory(inventory models.UpdateInventory, id int) (models.InventoryResponse, error) {
	query := `
	UPDATE inventories 
	SET stock = ?
	WHERE id = ?
	`

	if err := inv.DB.Exec(query, inventory.Stock, id).Error; err != nil {
		return models.InventoryResponse{}, err
	}

	var inventoryResponse models.InventoryResponse

	if err := inv.DB.Raw(`SELECT i.id AS product_id, i.product_name, i.category_id, c.category_name AS category, i.brand_id, b.brand_name AS brand, i.stock, i.price
  FROM inventories i
  INNER JOIN categories c ON i.category_id = c.id
  INNER JOIN brands b ON i.brand_id = b.id WHERE i.id = ?`, id).Scan(&inventoryResponse).Error; err != nil {
		return models.InventoryResponse{}, err
	}
	return inventoryResponse, nil
}

func (inv *InventoryRepostiory) ShowIndividualProduct(productID int) (models.InventoryResponse, error) {
	var inventory models.InventoryResponse

	query := `
	SELECT i.id AS product_id, i.product_name, i.category_id, c.category_name AS category, i.brand_id, b.brand_name AS brand, i.stock, i.price
  FROM inventories i
  INNER JOIN categories c ON i.category_id = c.id
  INNER JOIN brands b ON i.brand_id = b.id
	WHERE i.id = ?
  `

	if err := inv.DB.Raw(query, productID).Scan(&inventory).Error; err != nil {
		return models.InventoryResponse{}, err
	}

	return inventory, nil
}

func (inv *InventoryRepostiory) CheckStock(productID int) (models.CheckStockResponse, error) {
	var stockResponse models.CheckStockResponse

	query := `
	SELECT product_name,stock
	FROM inventories
	WHERE id = ?
	`
	if err := inv.DB.Raw(query, productID).Scan(&stockResponse).Error; err != nil {
		return models.CheckStockResponse{}, err
	}

	return stockResponse, nil
}

func (inv *InventoryRepostiory) ListProductsWithImages() ([]models.InventoryResponse, error) {
	var products []models.InventoryResponse

	query := `
        SELECT i.id AS product_id, i.product_name, i.category_id, c.category_name AS category, 
               i.brand_id, b.brand_name AS brand, i.stock, i.price
        FROM inventories i
        INNER JOIN categories c ON i.category_id = c.id
        INNER JOIN brands b ON i.brand_id = b.id
    `

	err := inv.DB.Raw(query).Scan(&products).Error
	if err != nil {
		return []models.InventoryResponse{}, err
	}
	return products, nil
}

func (inv *InventoryRepostiory) GetImages(productID int) ([]string,error) {
	query := `
	SELECT image FROM images WHERE inventory_id = ?
	`

	var images []string
	if err := inv.DB.Raw(query,productID).Scan(&images).Error; err != nil {
		return nil,err
	}

	return images,nil
}
