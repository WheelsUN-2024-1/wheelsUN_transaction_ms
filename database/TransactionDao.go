package database

import "gorm.io/gorm"

type TransactionDao struct {
	gorm.Model
	ReferenceCode  string `json:"referenceCode"`
	Description    string `json:"description"`
	Value          uint   `json:"value"`
	PaymentMethods string `json:"paymentMethods"`
	State          string `json:"state"`
	TripId         string    `json:"tripId"`
	CreditCardId   int    `json:"creditCardId"`
}

func (TransactionDao) TableName() string {
	return "transaction"
}
