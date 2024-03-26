package models

type PaymentMethod struct {
	Language string   `json:"language"`
	Command  string   `json:"command"`
	Test     bool     `json:"test"`
	Merchant Merchant `json:"merchant"`
}
