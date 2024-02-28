package routes

import (
	"github.com/ahdaan98/pkg/api/handler"
	"github.com/ahdaan98/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, categoryHandler *handler.CategoryHandler, brandHandler *handler.BrandHandler, inventoryHandler *handler.InventoryHandler, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler, walletHandler *handler.WalletHandler, couponHandler *handler.CouponHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.UserLogin)

	engine.POST("/verifyotp", userHandler.VerifyOTP)

	products:=engine.Group("/products")
	products.GET("/list",inventoryHandler.ListProductsWithImages)
	products.GET("", inventoryHandler.ListProducts)
	engine.GET("/categories/filter", categoryHandler.FilterByCategory)
	engine.GET("/brands/filter", brandHandler.FilterByBrand)
	products.GET("/filter/brand",brandHandler.FilterByBrand)

	engine.GET("/payment", paymentHandler.MakePaymentRazorpay) // Update this route
	engine.GET("/verifypayment", paymentHandler.VerifyPayment) // Update this route
	

	engine.GET("/invoice/print", orderHandler.PrintInvoice)

	engine.Use(middleware.UserAuthMiddleware)
	{

		profile := engine.Group("/profile")
		{
			profile.GET("", userHandler.UserProfile)
			profile.GET("/address", userHandler.GetAddress)
			profile.POST("", userHandler.AddAddress)
			profile.PUT("", userHandler.EditUserProfile)
			profile.PATCH("", userHandler.ChangePassword)

			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.GET("/all", orderHandler.GetAllOrders)
				orders.DELETE("", orderHandler.CancelOrder)
				orders.PUT("/return", orderHandler.ReturnOrder)
			}

		}

		cart := engine.Group("/cart")
		{
			cart.POST("", cartHandler.AddToCart)
			cart.GET("", userHandler.GetCart)
			cart.DELETE("", userHandler.RemoveFromCart)
			cart.PUT("", userHandler.UpdateQuantity)
		}

		checkout := engine.Group("/check-out")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.POST("", orderHandler.OrderItemsFromCart)
		}

		wallet := engine.Group("/wallet")
		{
			wallet.GET("", walletHandler.ViewWallet)

		}
		coupon := engine.Group("/coupon")
		{
			coupon.GET("", couponHandler.GetAllCoupons)
		}

	}
	
}
