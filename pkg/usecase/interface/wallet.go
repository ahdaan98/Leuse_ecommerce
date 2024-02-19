package interfaces

import "github.com/ahdaan98/pkg/utils/models"

type WalletUsecase interface {
	GetWallet(id int) (models.WalletAmount, error)
}