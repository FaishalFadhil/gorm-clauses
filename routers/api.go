package routers

import (
	"assignment-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders/:id", controllers.GetOneOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)
	router.GET("/items/:id", controllers.GetOneItem)
	router.GET("/items", controllers.GetAllItems)
	router.POST("/items", controllers.CreateItem)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)

	return router
}
