package models

type AdditionalInfo struct {
	PaymentNetwork                interface{} `json:"paymentNetwork"`
	RejectionType                 interface{} `json:"rejectionType"`
	ResponseNetworkMessage        interface{} `json:"responseNetworkMessage"`
	TravelAgencyAuthorizationCode interface{} `json:"travelAgencyAuthorizationCode"`
	CardType                      interface{} `json:"cardType"`
	TransactionType               interface{} `json:"transactionType"`
}
