package usecase

import (
	"errors"
	interfaces "github.com/ahdaan98/pkg/repository/interface"
	services "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userUseCase         services.UserUseCase
	adrepo              interfaces.AdminRepository
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userUseCase services.UserUseCase, adrepo interfaces.AdminRepository) services.CartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userUseCase:         userUseCase,
		adrepo:              adrepo,
	}
}

func (i *cartUseCase) AddToCart(userID, inventoryID, qty int) error {
	if userID < 1 || inventoryID < 1 || qty < 1 {
		return errors.New("check id properly, it cannot negative")
	}

	if userID == 0 || inventoryID == 0 || qty == 0 {
		return errors.New("check id properly, it cannot zero")
	}
	
	cart_id, err := i.repo.GetCartId(userID)
	if err != nil {
		return errors.New("some error in geting user cart")
	}

	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(userID)
		if err != nil {
			return errors.New("cannot create cart fro user")
		}
	}

	exists, err := i.repo.CheckIfItemIsAlreadyAdded(cart_id, inventoryID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("item already exists in cart")
	}

	if err := i.repo.AddLineItems(cart_id, inventoryID, qty); err != nil {
		return errors.New("error in adding products")
	}

	return nil
}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	paymethods, err := i.adrepo.GetPaymentMethod()
	if err != nil {
		return models.CheckOut{}, err
	}

	products, err := i.userUseCase.GetCart(id)
	if err != nil {
		return models.CheckOut{}, err
	}
	var checkout models.CheckOut

	checkout.CartID = products.ID
	checkout.Addresses = address
	checkout.Products = products.Data
	checkout.PaymentMethod = paymethods

	return checkout, err
}
