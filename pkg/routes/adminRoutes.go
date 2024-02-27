package routes

import (
	"github.com/ahdaan98/pkg/api/handler"
	"github.com/ahdaan98/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, categoryHandler *handler.CategoryHandler, brandHandler *handler.BrandHandler, inventoryHandler *handler.InventoryHandler, adminHandler *handler.AdminHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler) {

	engine.POST("/login", adminHandler.AdminLogin)
	engine.Static("/uploads", "./uploads")

	engine.GET("/sales-by-date", adminHandler.SalesByDate)
	engine.GET("/custom-report", adminHandler.CustomSalesReport)
	engine.GET("/sales/excel/download", handler.DownloadExcel)

	engine.Use(middleware.AuthMiddleware)
	{
		user := engine.Group("/users")
		{
			user.GET("", adminHandler.GetUsers)
			user.GET("/:id", adminHandler.GetUserByID)
			user.PUT("/block", adminHandler.BlockUser)
			user.PUT("/unblock", adminHandler.UnBlockUser)
		}
		inventory := engine.Group("/products")
		{
			inventory.GET("", inventoryHandler.ListProducts)
			inventory.GET("/list",inventoryHandler.ListProductsWithImages)
			inventory.POST("/add", inventoryHandler.AddInventory)
			inventory.POST("/image", inventoryHandler.UploadProductImage)
			inventory.PUT("/edit", inventoryHandler.EditInventory)
			inventory.PUT("/update/stock", inventoryHandler.UpdateInventory)
			inventory.GET("/stock", inventoryHandler.CheckStock)
			inventory.GET("/:id", inventoryHandler.ShowIndividualProduct)
		}

		category := engine.Group("/categories")
		{
			category.GET("", categoryHandler.GetCategories)
			category.POST("/add", categoryHandler.AddCategory)
			category.PUT("/edit", categoryHandler.EditCategory)
			category.DELETE("/:id", categoryHandler.DeleteCategory)
			category.GET("/filter", categoryHandler.FilterByCategory)
		}

		brand := engine.Group("/brands")
		{
			brand.GET("", brandHandler.GetBrands)
			brand.POST("/add", brandHandler.AddBrand)
			brand.PUT("/edit", brandHandler.EditBrand)
			brand.DELETE("/:id", brandHandler.DeleteBrand)
			brand.GET("/filter", brandHandler.FilterByBrand)
		}

		payment := engine.Group("/payment-methods")
		{
			payment.POST("", adminHandler.NewPaymentMethod)
			payment.GET("", adminHandler.ListPaymentMethods)
			payment.DELETE("/:id", adminHandler.DeletePaymentMethod)
		}

		orders := engine.Group("/orders")
		{
			orders.GET("", orderHandler.GetAdminOrders)
			orders.PUT("/status", orderHandler.ApproveOrder)
		}

		coupon := engine.Group("/coupons")
		{
			coupon.POST("", couponHandler.AddCoupon)
			coupon.GET("", couponHandler.GetCoupons)
			coupon.PATCH("/:id", couponHandler.UpdateCoupon)
		}
		engine.GET("/dashboard", adminHandler.DashBoard)
	}
}
