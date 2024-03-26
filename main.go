package main

import (
	"wheelsUN_transaction_ms/configs"
	"wheelsUN_transaction_ms/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}
func main() {
	// Create a new Gin engine
	router := gin.Default()

	// Configure routes

	routes.TransactionRouter(router)
	routes.CreditCardRouter(router)

	router.Run()
}
