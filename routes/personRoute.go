package routes

import (
	"wheelsUN_transaction_ms/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRouter(router *gin.Engine) {

	routes := router.Group("api/persons")
	routes.POST("", controllers.PersonCreate)
	routes.GET("", controllers.PersonGet)
	routes.GET("/:id", controllers.PersonGetById)
	routes.PUT("/:id", controllers.PersonUpdate)
	routes.DELETE("/:id", controllers.PersonDelete)

}
