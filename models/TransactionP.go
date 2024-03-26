package models

type TransactionP struct {
	Order           Order           `json:"order"`
	CreditCard      CreditCard      `json:"creditCard"`
	Payer           Payer           `json:"payer"`
	Type            string          `json:"type"`
	PaymentMethod   string          `json:"paymentMethod"`
	PaymentCountry  string          `json:"paymentCountry"`
	DeviceSessionID string          `json:"deviceSessionId"`
	IPAddress       string          `json:"ipAddress"`
	Cookie          string          `json:"cookie"`
	UserAgent       string          `json:"userAgent"`
	ExtraParameters ExtraParameters `json:"extraParameters"`
}
