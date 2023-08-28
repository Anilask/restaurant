package routes

import (
	"restaurant/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.Getorders())
	incomingRoutes.GET("/orders/:orders_id", controller.Getorder())
	incomingRoutes.POST("/orders", controller.Createorder())
	incomingRoutes.PATCH("/orders/:orders_id", controller.Updateorder())
}
