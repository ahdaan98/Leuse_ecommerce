package interfaces

import "github.com/ahdaan98/pkg/utils/models"

type WalletRepository interface {
	GetWallet(userID int) (models.WalletAmount, error)
	//GetsWallet(orderId int) (models.WalletAmount, error)
	AddToWallet(Price, UserId int) (models.WalletAmount, error)
}