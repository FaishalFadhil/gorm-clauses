package routers

import (
	"clauses/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders/:id", controllers.GetOneOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/clause-normal/:id", controllers.UpdateOrder)
	router.PUT("/orders/clause-doNothing/:id", controllers.UpdateDoNothingOrder)
	router.PUT("/orders/clause-noConflict/:id", controllers.UpdateNoConflictOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	return router
}
