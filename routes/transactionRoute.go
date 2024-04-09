package routes

import (
	"wheelsUN_transaction_ms/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(router *gin.Engine) {
	routes := router.Group("/transaction")
	{
		routes.POST("/create", func(c *gin.Context) {
			// Envolver la función del controlador dentro de gin.HandlerFunc
			controllers.PostCardPayment(c.Writer, c.Request)
		})

		routes.POST("/createdata", func(c *gin.Context) {
			// Envolver la función del controlador dentro de gin.HandlerFunc
			controllers.PostInDatabase(c.Writer, c.Request)
		})

		routes.GET("/get/:referenceCode", func(c *gin.Context) {
			// Obtén el código de referencia de la URL utilizando c.Param("referenceCode")
			referenceCode := c.Param("referenceCode")
			// Llama al método GetTransactionReferenceCode del controlador
			controllers.GetTransactionReferenceCode(c.Writer, c.Request, referenceCode)
		})

	}
}
