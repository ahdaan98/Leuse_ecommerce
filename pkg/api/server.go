package http

import (
	"log"

	"github.com/ahdaan98/pkg/api/handler"
	"github.com/ahdaan98/pkg/config"
	"github.com/ahdaan98/pkg/routes"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler,brandHandler *handler.BrandHandler, inventoryHandler *handler.InventoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler, walletHandler *handler.WalletHandler, couponHandler *handler.CouponHandler) *ServerHTTP {
	engine := gin.Default()

	engine.LoadHTMLGlob("pkg/templates/*.html")
	routes.UserRoutes(engine.Group("/user"), categoryHandler, brandHandler, inventoryHandler, userHandler, cartHandler, orderHandler, paymentHandler, walletHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/admin"),categoryHandler, brandHandler, inventoryHandler,adminHandler,orderHandler, couponHandler)

	return &ServerHTTP{
		engine: engine,
	}
}
func (s *ServerHTTP) Start() {
	cfg,_:=config.LoadEnvVariables()
	err := s.engine.Run(":"+cfg.PORT)
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
}
