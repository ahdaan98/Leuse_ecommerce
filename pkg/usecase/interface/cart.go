package interfaces

import "github.com/ahdaan98/pkg/utils/models"

type CartUseCase interface {
	AddToCart(user_id, inventory_id, qty int) error
	CheckOut(id int) (models.CheckOut, error)
}