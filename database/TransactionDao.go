package database

import "gorm.io/gorm"

type TransactionDao struct {
	gorm.Model
	TransactionId    string `json:"transactionId"`
	ReferenceCode    string `json:"referenceCode"`
	Description      string `json:"description"`
	Value            uint   `json:"value"`
	PaymentMethods   string `json:"paymentMethods"`
	State            string `json:"state"`
	TransactionIdPay string `json:"transactionIdPay"`
	OrderId          string `json:"orderId"`
	TripId           int    `json:"tripId"`
	CreditCardId     int    `json:"creditCardId"`
}

func (TransactionDao) TableName() string {
	return "transaction"
}
