package routes

import (
	"wheelsUN_transaction_ms/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(router *gin.Engine) {
	routes := router.Group("/api/transactions")
	{
		routes.POST("cardpayment", func(c *gin.Context) {
			// Envolver la función del controlador dentro de gin.HandlerFunc
			controllers.PostCardPayment(c.Writer, c.Request)
		})
		// Agrega otras rutas relacionadas con las transacciones si es necesario
	}
}
