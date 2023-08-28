package routes

import (
	"restaurant/controller"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.Gettables())
	incomingRoutes.GET("/tables/:tables_id", controller.Gettable())
	incomingRoutes.POST("/tables", controller.Createtable())
	incomingRoutes.PATCH("/tables/:tables_id", controller.Updatetable())
}
