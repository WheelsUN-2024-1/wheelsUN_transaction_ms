package models

type Order struct {
	AccountID        string           `json:"accountId"`
	ReferenceCode    string           `json:"referenceCode"`
	Description      string           `json:"description"`
	Language         string           `json:"language"`
	NotifyURL        string           `json:"notifyUrl"`
	PartnerID        string           `json:"partnerId"`
	Signature        string           `json:"signature"`
	ShippingAddress  Address          `json:"shippingAddress"`
	Buyer            Buyer            `json:"buyer"`
	AdditionalValues AdditionalValues `json:"additionalValues"`
}
