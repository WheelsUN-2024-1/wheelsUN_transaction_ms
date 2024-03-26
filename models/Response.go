package models

type Response struct {
	Code                string              `json:"code"`
	Error               interface{}         `json:"error"`
	TransactionResponse TransactionResponse `json:"transactionResponse"`
}

type TransactionResponse struct {
	Order                       int64                 `json:"orderId"`
	TransactionId               string                `json:"transactionId"`
	State                       string                `json:"state"`
	PaymentNetworkResponseCode  string                `json:"paymentNetworkResponseCode"`
	PaymentNetworkResponseError interface{}           `json:"paymentNetworkResponseErrorMessage"`
	TrazabilityCode             string                `json:"trazabilityCode"`
	AuthorizationCode           string                `json:"authorizationCode"`
	PendingReason               interface{}           `json:"pendingReason"`
	ResponseCode                string                `json:"responseCode"`
	ErrorCode                   interface{}           `json:"errorCode"`
	ResponseMessage             string                `json:"responseMessage"`
	TransactionDate             interface{}           `json:"transactionDate"`
	TransactionTime             interface{}           `json:"transactionTime"`
	OperationDate               int64                 `json:"operationDate"`
	ReferenceQuestionnaire      interface{}           `json:"referenceQuestionnaire"`
	ExtraParameters             map[string]string     `json:"extraParameters"`
	AdditionalInfo              TransactionAdditional `json:"additionalInfo"`
}

type TransactionAdditional struct {
	PaymentNetwork            string `json:"paymentNetwork"`
	RejectionType             string `json:"rejectionType"`
	ResponseNetworkMessage    string `json:"responseNetworkMessage"`
	TravelAgencyAuthorization string `json:"travelAgencyAuthorizationCode"`
	CardType                  string `json:"cardType"`
	TransactionType           string `json:"transactionType"`
}
