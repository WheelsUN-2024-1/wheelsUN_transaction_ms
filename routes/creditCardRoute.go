package routes

import (
	"wheelsUN_transaction_ms/controllers"

	"github.com/gin-gonic/gin"
)

func CreditCardRouter(router *gin.Engine) {
	routes := router.Group("/api/creditcard")
	{
		routes.POST("/create", func(c *gin.Context) {
			// Envolver la función del controlador dentro de gin.HandlerFunc
			controllers.PostCreditCard(c.Writer, c.Request)
		})

		routes.GET("/get", func(c *gin.Context) {
			// Llama al método GetCreditCardByID del controlador
			controllers.GetCreditCardByID(c.Writer, c.Request)
		})

		routes.PUT("/put", func(c *gin.Context) {
			// Llama al método GetCreditCardByID del controlador
			controllers.PutCreditCard(c.Writer, c.Request)
		})
	}
}
