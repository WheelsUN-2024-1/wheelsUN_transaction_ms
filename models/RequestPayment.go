package models

type RequestPayment struct {
	Language    string       `json:"language"`
	Command     string       `json:"command"`
	Test        bool         `json:"test"`
	Merchant    Merchant     `json:"merchant"`
	Transaction TransactionP `json:"transaction"`
}
