// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ahdaan98/pkg/api"
	"github.com/ahdaan98/pkg/api/handler"
	"github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/db"
	"github.com/ahdaan98/pkg/helper"
	"github.com/ahdaan98/pkg/repository"
	"github.com/ahdaan98/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDB(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	interfacesHelper := helper.NewHelper(cfg)
	inventoryRepository := repository.NewInventoryRespository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, interfacesHelper, cfg, inventoryRepository)
	userHandler := handler.NewUserHandler(userUseCase, interfacesHelper)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, userRepository,interfacesHelper)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository,interfacesHelper)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)
	brandRepository := repository.NewBrandRepository(gormDB)
	brandUseCase := usecase.NewBrandUseCase(brandRepository,interfacesHelper)
	brandHandler := handler.NewBrandHandler(brandUseCase)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)
	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, inventoryRepository, userUseCase, adminRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)
	orderRepository := repository.NewOrderRepository(gormDB)
	walletRepository := repository.NewWalletRepository(gormDB)
	couponRepository := repository.NewCouponRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, userUseCase, walletRepository, cartRepository, couponRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(orderRepository, paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)
	walletUsecase := usecase.NewWalletUseCase(walletRepository)
	walletHandler := handler.NewWalletHandler(walletUsecase)
	couponUseCase := usecase.NewCouponUseCase(couponRepository, orderRepository, cartRepository)
	couponHandler := handler.NewCouponHandler(couponUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, categoryHandler, brandHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler, walletHandler, couponHandler)
	return serverHTTP, nil
}
