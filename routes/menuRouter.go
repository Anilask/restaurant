package routes

import (
	"restaurant/controller"

	"github.com/gin-gonic/gin"
)

func MenuRoutesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.Getmenus())
	incomingRoutes.GET("/menus/:menus_id", controller.Getmenu())
	incomingRoutes.POST("/menus", controller.Createmenu())
	incomingRoutes.PATCH("/menus/:menus_id", controller.Updatemenu())
}
