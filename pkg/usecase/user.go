package usecase

import (
	"errors"

	"github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/domain"
	helper "github.com/ahdaan98/pkg/helper/interfaces"
	repo "github.com/ahdaan98/pkg/repository/interface"
	service "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
)

type UserUseCase struct {
	repo                repo.UserRepository
	helper              helper.Helper
	cfg                 config.Config
	inventoryRepository repo.InventoryRepository
}

func NewUserUseCase(repo repo.UserRepository, helper helper.Helper, cfg config.Config, inv repo.InventoryRepository) service.UserUseCase {
	return &UserUseCase{
		repo:                repo,
		helper:              helper,
		cfg:                 cfg,
		inventoryRepository: inv,
	}
}

var InternalError = "Internal Server Error"
var ErrorHashingPassword = "Error In Hashing Password"

func (u *UserUseCase) ValidatingDetails(user models.UserSignUp) error {
	if err := u.helper.ValidateName(user.Name); err != nil {
		return err
	}

	// Validate user's email
	if err := u.helper.ValidateEmail(user.Email); err != nil {
		return err
	}

	// Check if user with the provided email already exists
	exist, err := u.repo.CheckUserExist(user.Email)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user with this email already exists")
	}

	// Validate user's phone number
	if !u.helper.ValidatePhoneNumber(user.Phone) {
		return errors.New("invalid phone number")
	}

	// Validate user's password
	if err := u.helper.ValidatePassword(user.Password, user.ConfirmPassword); err != nil {
		return err
	}

	return nil

}

func (u *UserUseCase) UserSignUp(user models.UserSignUp) (models.TokenUsers, error) {
	// Create user
	resp, err := u.repo.CreateUser(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// Generate token
	token, err := u.helper.GenerateTokenClients(resp)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: resp,
		Token: token,
	}, nil
}

func (u *UserUseCase) UserLogin(user models.UserLogin) (models.TokenUsers, error) {
	// Check if the user exists
	exists, err := u.repo.CheckUserExist(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}
	if !exists {
		return models.TokenUsers{}, errors.New("user does not exist")
	}

	// Retrieve user data
	resp, err := u.repo.GetUserByEmail(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// Retrieve user's password
	storedPassword, err := u.repo.GetUserPassword(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// Verify password
	if storedPassword != user.Password {
		return models.TokenUsers{}, errors.New("incorrect password")
	}

	// Check if the user is blocked
	blocked, err := u.repo.CheckBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}
	if blocked {
		return models.TokenUsers{}, errors.New("you are blocked by admin")
	}

	// Generate token
	token, err := u.helper.GenerateTokenClients(resp)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: resp,
		Token: token,
	}, nil
}

func (u *UserUseCase) UserProfile(id int) (models.UserDetailsResponse, error) {
	resp, err := u.repo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return resp, nil
}

func (u *UserUseCase) EditProfile(details models.EditUserDetails, id int) (models.UserDetailsResponse, error) {
	resp, err := u.repo.EditDetails(details, id)
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return resp, nil
}

func (u *UserUseCase) AddAddress(id int, address models.AddAddress) error {
	err := u.repo.AddAddress(id, address)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) GetAddresses(id int) ([]domain.Address, error) {
	addresses, err := u.repo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil
}

func (u *UserUseCase) GetCart(id int) (models.GetCartResponse, error) {

	//find cart id
	cart_id, err := u.repo.GetCartID(id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	//find products inide cart
	products, err := u.repo.GetProductsInCart(cart_id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	//find product names
	var product_names []string
	for i := range products {
		product_name, err := u.repo.FindProductNames(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		product_names = append(product_names, product_name)
	}

	//find quantity
	var quantity []int
	for i := range products {
		q, err := u.repo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		quantity = append(quantity, q)
	}

	var price []float64
	for i := range products {
		q, err := u.repo.FindPrice(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		price = append(price, q)
	}

	var categoryID []int
	for i := range products {
		c, err := u.repo.FindCategory(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		categoryID = append(categoryID, c)
	}

	var brandID []int
	for i := range products {
		c, err := u.repo.FindBrand(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		brandID = append(brandID, c)
	}

	var category []string
	for i := range products {
		c, err := u.repo.FindCategoryName(categoryID[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		category = append(category, c)
	}

	var brand []string
	for i := range products {
		c, err := u.repo.FindBrandName(brandID[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		brand = append(brand, c)
	}

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ProductID = products[i]
		get.ProductName = product_names[i]
		get.BrandID = uint(brandID[i])
		get.Brand = brand[i]
		get.CategoryID = uint(categoryID[i])
		get.Category = category[i]
		get.Quantity = quantity[i]
		get.Price = int(price[i])
		get.Total = (price[i]) * float64(quantity[i])

		getcart = append(getcart, get)
	}

	var response models.GetCartResponse
	response.ID = cart_id
	response.Data = getcart
	//then return in appropriate format

	return response, nil
}

func (i *UserUseCase) RemoveFromCart(userID, inventory int) error {

	exist,err:=i.inventoryRepository.CheckInventoryExistByID(inventory)
	if err!=nil{
		return err
	}

	if !exist {
		return errors.New("inventory with id does not exist")
	} 

	cart, err := i.repo.GetCartID(userID)
	if err != nil {
		return err
	}

	err = i.repo.RemoveFromCart(cart, inventory)
	if err != nil {
		return err
	}

	return nil

}

func (i *UserUseCase) UpdateQuantity(id, inv, qty int) error {

	if id <= 0 || inv <= 0 || qty <=0 {
		return errors.New("check values properly, values cannot be negative or zero")
	}

	stock, err := i.inventoryRepository.CheckStock(inv)
	if err != nil {
		return err
	}

	if qty > stock.Stock {
		return errors.New("out of stock")
	}

	err = i.repo.UpdateQuantity(id, inv, qty)
	return err
}

func (i *UserUseCase) ChangePassword(id int, old string, password string, repassword string) error {

	if id <= 0 {
		return errors.New("invalid id")
	}

	user, err := i.repo.GetUserDetails(id)
	if err != nil {
		return err
	}

	userPassword, err := i.repo.GetUserPassword(user.Email)
	if err != nil {
		return err
	}

	if old != userPassword {
		return errors.New("please enter correct password")
	}
	if password != repassword {
		return errors.New("passwords does not match")
	}

	err=i.helper.ValidatePassword(password,repassword)
	if err!=nil{
		return err
	}

	return i.repo.ChangePassword(id, string(password))

}

func (i *UserUseCase) CheckUserExistByEmail(email string) (bool, error) {
	return i.repo.CheckUserExist(email)
}
