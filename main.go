package main

import (
	"database/sql"
	"os"
	"restaurant/middleware"
	"restaurant/routes"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRouter(router)
	router.Use(middleware.Authentication)

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OderItemRoutes(router)
	routes.InvoicesRoutes(router)

	router.Run(":" + port)
}
