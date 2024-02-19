package usecase

import (
	"github.com/ahdaan98/pkg/repository/interface"
	services  "github.com/ahdaan98/pkg/usecase/interface"

	"github.com/ahdaan98/pkg/utils/models"
)

type walletUseCase struct {
	walletRepository interfaces.WalletRepository
}

func NewWalletUseCase(repository interfaces.WalletRepository) services.WalletUsecase {
	return &walletUseCase{
		walletRepository: repository,
	}
}

func (wt *walletUseCase) GetWallet(userID int) (models.WalletAmount, error) {
	return wt.walletRepository.GetWallet(userID)
}