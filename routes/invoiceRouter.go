package routes

import (
	"restaurant/controller"

	"github.com/gin-gonic/gin"
)

func InvoicesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoices_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoices_id", controller.UpdateInvoice())
}
