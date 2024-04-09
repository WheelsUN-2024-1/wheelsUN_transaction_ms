package models

type ResponseJSON struct {
	Code                string                  `json:"code"`
	Error               interface{}             `json:"error,omitempty"`
	TransactionResponse TransactionResponseJSON `json:"transactionResponse"`
}

type TransactionResponseJSON struct {
	Order                       string                    `json:"orderId"`
	TransactionId               string                    `json:"transactionId"`
	State                       string                    `json:"state"`
	PaymentNetworkResponseCode  string                    `json:"paymentNetworkResponseCode"`
	PaymentNetworkResponseError interface{}               `json:"paymentNetworkResponseErrorMessage,omitempty"`
	TrazabilityCode             string                    `json:"trazabilityCode"`
	AuthorizationCode           string                    `json:"authorizationCode"`
	PendingReason               interface{}               `json:"pendingReason,omitempty"`
	ResponseCode                string                    `json:"responseCode"`
	ErrorCode                   interface{}               `json:"errorCode,omitempty"`
	ResponseMessage             string                    `json:"responseMessage"`
	TransactionDate             interface{}               `json:"transactionDate,omitempty"`
	TransactionTime             interface{}               `json:"transactionTime,omitempty"`
	OperationDate               string                    `json:"operationDate"`
	ReferenceQuestionnaire      interface{}               `json:"referenceQuestionnaire,omitempty"`
	ExtraParameters             ExtraParametersJSON       `json:"extraParameters"`
	AdditionalInfo              TransactionAdditionalJSON `json:"additionalInfo"`
}

type ExtraParametersJSON struct {
	Entries []ExtraParametersEntryJSON `json:"entry"`
}

type ExtraParametersEntryJSON struct {
	Key   string `json:"string"`
	Value string `json:"string"`
}

type TransactionAdditionalJSON struct {
	PaymentNetwork            string `json:"paymentNetwork"`
	RejectionType             string `json:"rejectionType"`
	ResponseNetworkMessage    string `json:"responseNetworkMessage"`
	TravelAgencyAuthorization string `json:"travelAgencyAuthorization,omitempty"`
	CardType                  string `json:"cardType"`
	TransactionType           string `json:"transactionType"`
}
