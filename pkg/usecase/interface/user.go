package interfaces

import (
	"github.com/ahdaan98/pkg/domain"
	"github.com/ahdaan98/pkg/utils/models"
)

type UserUseCase interface {
	ValidatingDetails(models.UserSignUp) error
	UserSignUp(user models.UserSignUp) (models.TokenUsers, error)
	UserLogin(user models.UserLogin) (models.TokenUsers, error)
	UserProfile(id int) (models.UserDetailsResponse, error)
	EditProfile(details models.EditUserDetails, id int) (models.UserDetailsResponse, error)

	AddAddress(id int, address models.AddAddress) error
	GetAddresses(id int) ([]domain.Address, error)

	GetCart(id int) (models.GetCartResponse, error)
	RemoveFromCart(userID, inventory int) error
	UpdateQuantity(id, inv_id, qty int) error

	ChangePassword(id int, old string, password string, repassword string) error
	CheckUserExistByEmail(email string) (bool, error)
}
