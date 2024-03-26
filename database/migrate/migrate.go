package main

import (
	"wheelsUN_transaction_ms/configs"
	"wheelsUN_transaction_ms/database"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&database.CreditCardDao{}, &database.TransactionDao{})
}
