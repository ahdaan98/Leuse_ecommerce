package usecase

import (
	"errors"

	"github.com/ahdaan98/pkg/config"
	repo "github.com/ahdaan98/pkg/repository/interface"
	usecase "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type InventoryUseCase struct {
	repository repo.InventoryRepository
}

func NewInventoryUseCase(repo repo.InventoryRepository) usecase.InventoryUseCase {
	return &InventoryUseCase{
		repository: repo,
	}
}

func (i *InventoryUseCase) AddInventory(inventory models.AddInventory) (models.InventoryResponse, error) {
	Exist, err := i.repository.CheckInventoryExist(inventory.ProductName)
	if err != nil || Exist {
		return models.InventoryResponse{}, err
	}

	k, err := i.repository.AddInventory(inventory)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return k, nil
}

func (i *InventoryUseCase) ListProducts(page, per_product int) ([]models.InventoryResponse, error) {
	k, err := i.repository.ListProducts(page, per_product)
	if err != nil {
		return []models.InventoryResponse{}, err
	}
	return k, nil
}

func (i *InventoryUseCase) EditInventory(inventory models.EditInventory, id int) (models.InventoryResponse, error) {
	Exist, err := i.repository.CheckInventoryExist(inventory.ProductName)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	if Exist {
		return models.InventoryResponse{}, errors.New("product already exist")
	}

	Exist1, err1 := i.repository.CheckInventoryExistByID(id)
	if err1 != nil {
		return models.InventoryResponse{}, err
	}

	if !Exist1 {
		return models.InventoryResponse{}, errors.New("product does not exist with this id")
	}

	inv, err := i.repository.EditInventory(inventory, id)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return inv, nil
}

func (i *InventoryUseCase) UpdateInventory(inventory models.UpdateInventory, id int) (models.InventoryResponse, error) {
	Exist, err := i.repository.CheckInventoryExistByID(id)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	if !Exist {
		return models.InventoryResponse{}, errors.New("product does not exist with this id")
	}

	inv, err := i.repository.UpdateInventory(inventory, id)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return inv, nil
}

func (i *InventoryUseCase) ShowIndividualProduct(productID int) (models.InventoryResponse, error) {
	Exist, err := i.repository.CheckInventoryExistByID(productID)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	if !Exist {
		return models.InventoryResponse{}, errors.New("product does not exist with this id")
	}

	inv, err := i.repository.ShowIndividualProduct(productID)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return inv, nil
}

func (i *InventoryUseCase) CheckStock(productID int) (models.CheckStockResponse, error) {
	Exist, err := i.repository.CheckInventoryExistByID(productID)
	if err != nil {
		return models.CheckStockResponse{}, err
	}

	if !Exist {
		return models.CheckStockResponse{}, errors.New("product does not exist with this id")
	}

	inv, err := i.repository.CheckStock(productID)
	if err != nil {
		return models.CheckStockResponse{}, err
	}

	return inv, nil
}

func (i *InventoryUseCase) AddImage(id int, image string) error {
	return i.repository.UploadImage(id, image)
}

func (uc *InventoryUseCase) ListProductsWithImages() ([]models.InventoryResponseWithImages, error) {
	cfg,_:=config.LoadEnvVariables()
    productList, err := uc.repository.ListProductsWithImages()
    if err != nil {
        return nil, err
    }
    var responseList []models.InventoryResponseWithImages
    // Iterate over the product list and collect images
    for _, product := range productList {
        images, err := uc.repository.GetImages(int(product.ProductID))
        if err != nil {
            return nil, err
        }

        var urls []string
        for _, image := range images {
            url := "http://localhost:"+cfg.PORT+"/admin/uploads/" + image
            urls = append(urls, url)
        }

        response := models.InventoryResponseWithImages{
            ProductID:   product.ProductID,
            ProductName: product.ProductName,
            CategoryID:  product.CategoryID,
            Category:    product.Category,
            BrandID:     product.BrandID,
            Brand:       product.Brand,
            Stock:       product.Stock,
            Price:       product.Price,
            Images:      urls,
        }
        responseList = append(responseList, response)
    }

    return responseList, nil
}