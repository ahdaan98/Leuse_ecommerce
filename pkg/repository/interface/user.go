package interfaces

import (
	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type UserRepository interface {
	CreateUser(user models.UserSignUp) (models.UserDetailsResponse,error)

	CheckUserExist(email string) (bool,error)
	CheckUserExistByID(id int) (bool,error)
	CheckBlockStatus(email string) (bool,error)

	GetUserDetails(id int) (models.UserDetailsResponse,error)
	EditDetails(details models.EditUserDetails,id int) (models.UserDetailsResponse,error)

	GetUserByEmail(email string) (models.UserDetailsResponse,error)
	GetUserPassword(email string) (string,error)

	AddAddress(userID int, address models.AddAddress) error
	GetAddresses(id int) ([]domain.Address, error)
	CheckIfFirstAddress(id int) bool

	ChangePassword(id int, password string) error

	GetCartID(id int) (int, error)
	GetProductsInCart(cart_id int) ([]int, error)
	FindProductNames(inventory_id int) (string, error)
	FindCartQuantity(cart_id, inventory_id int) (int, error)
	FindPrice(inventory_id int) (float64, error)
	FindStock(id int) (int, error)
	FindCategory(inventory_id int) (int, error)
	FindBrand(inventory_id int) (int, error)
	FindCategoryName(category_id int) (string, error)
	FindBrandName(brand_id int) (string, error)

	RemoveFromCart(cart, inventory int) error
	UpdateQuantity(id, inv_id, qty int) error
}