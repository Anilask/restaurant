package routes

import (
    "github.com/gin-gonic/gin"
    "restaurant/controller"
)
func UserRouter(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.Singup())
	incomingRoutes.POST("/users/login", controller.Login())
}