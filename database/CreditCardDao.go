package database

import "gorm.io/gorm"

type CreditCardDao struct {
	gorm.Model
	CreditCardId   uint   `json:"creditCardId"`
	UserId         int    `json:"userId"`
	Number         string `json:"number"`
	Name           string `json:"name"`
	SecurityCode   string `json:"securityCode"`
	ExpirationDate string `json:"expirationDate"`
}

func (CreditCardDao) TableName() string {
	return "creditCard"
}
