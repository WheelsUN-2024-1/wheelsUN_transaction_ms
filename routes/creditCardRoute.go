package routes

import (
	"wheelsUN_transaction_ms/controllers"

	"github.com/gin-gonic/gin"
)

func CreditCardRouter(router *gin.Engine) {
	routes := router.Group("/creditcard")
	{
		routes.POST("/create", func(c *gin.Context) {
			// Envolver la función del controlador dentro de gin.HandlerFunc
			controllers.PostCreditCard(c.Writer, c.Request)
		})

		routes.GET("/get/:id", func(c *gin.Context) {
			// Obtén el ID de la URL utilizando c.Param("id")
			id := c.Param("id")
			// Llama al método GetCreditCardByID del controlador
			controllers.GetCreditCardByID(c.Writer, id)
		})
		routes.GET("/cards/:id", func(c *gin.Context) {
			// Obtén el ID de la URL utilizando c.Param("id")
			id := c.Param("id")
			// Llama al método GetCreditCardByID del controlador
			controllers.GetCreditCardsByUserID(c.Writer, id)
		})

		routes.PUT("/put/:id", func(c *gin.Context) {
			// Obtén el ID de la URL utilizando c.Param("id")
			id := c.Param("id")
			// Llama al método PutCreditCard del controlador
			controllers.PutCreditCard(c.Writer, c.Request, id)
		})

		routes.DELETE("/delete/:id", func(c *gin.Context) {
			// Obtén el ID de la URL utilizando c.Param("id")
			id := c.Param("id")
			// Llama al método PutCreditCard del controlador
			controllers.DeleteCreditCard(c.Writer, id)
		})
	}
}
