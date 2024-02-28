package interfaces

import (
	"github.com/ahdaan98/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.AddInventory) (models.InventoryResponse, error)
	CheckInventoryExist(productName string) (bool, error)
	CheckInventoryExistByID(id int) (bool, error)
	ListProducts(page, per_product int) ([]models.InventoryResponse, error)
	EditInventory(inventory models.EditInventory, id int) (models.InventoryResponse, error)
	UpdateInventory(inventory models.UpdateInventory, id int) (models.InventoryResponse, error)
	ShowIndividualProduct(productID int) (models.InventoryResponse, error)
	CheckStock(productID int) (models.CheckStockResponse, error)

	UploadImage(id int, image string) error
	ListProductsWithImages(page, per_product int) ([]models.InventoryResponse, error)
	GetImages(productID int) ([]string,error) 
}
