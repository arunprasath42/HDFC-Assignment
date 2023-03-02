// This packege demonstrates routing with Gin framework
package route

import (
	"hdfc-assignment/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	//endpoint for productService
	productGroup := router.Group("/products")
	{
		productGroup.GET("", controllers.GetAllProducts)
		productGroup.GET("/details", controllers.GetProductDetails)
		productGroup.POST("/placeorder", controllers.PlaceOrder)
	}

	// endpoints for the Orderservice.
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/orderDetails", controllers.GetOrderDetails)
		orderGroup.POST("/updateOrderStatus", controllers.UpdateOrderStatus)
	}
	router.Run(viper.GetString("server.port"))
}
