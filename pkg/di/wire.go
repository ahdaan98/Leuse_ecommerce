//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/ahdaan98/pkg/api"
	"github.com/ahdaan98/pkg/api/handler"
	"github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/db"
	helper "github.com/ahdaan98/pkg/helper"
	"github.com/ahdaan98/pkg/repository"
	"github.com/ahdaan98/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP,error){
	 wire.Build(
		db.ConnectDB,

		helper.NewHelper,

		handler.NewBrandHandler,
		handler.NewCategoryHandler,
		handler.NewInventoryHandler,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewPaymentHandler,
		handler.NewWalletHandler,
		handler.NewCouponHandler,

		usecase.NewBrandUseCase,
		usecase.NewCategoryUseCase,
		usecase.NewInventoryUseCase,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewCartUseCase,
		usecase.NewOrderUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewWalletUseCase,
		usecase.NewCouponUseCase,

		repository.NewBrandRepository,
		repository.NewCategoryRepository,
		repository.NewInventoryRespository,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepository,
		repository.NewWalletRepository,
		repository.NewCouponRepository,

		http.NewServerHTTP,
	 )

	 return &http.ServerHTTP{},nil
}