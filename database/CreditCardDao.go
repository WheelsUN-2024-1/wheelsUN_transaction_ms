package database

import "gorm.io/gorm"

type CreditCardDao struct {
	gorm.Model
	CreditCardId   uint   `json:"creditCardId"`
	UserId         string    `json:"userId"`
	Number         string `json:"number"`
	Name           string `json:"name"`
	SecurityCode   string `json:"securityCode"`
	ExpirationDate string `json:"expirationDate"`
	Brand          string `json:"brand"`
}

func (CreditCardDao) TableName() string {
	return "creditCard"
}
