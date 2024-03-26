package models

type CreditCard struct {
	Number             string `json:"number"`
	SecurityCode       string    `json:"securityCode"`
	ExpirationDate     string `json:"expirationDate"`
	Name               string `json:"name"`
	ProcessWithoutCvv2 bool   `json:"processWithoutCvv2"`
}
