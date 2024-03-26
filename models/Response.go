package models

import "encoding/xml"

type Response struct {
	XMLName             xml.Name            `xml:"paymentResponse"`
	Code                string              `xml:"code"`
	Error               interface{}         `xml:"error,omitempty"`
	TransactionResponse TransactionResponse `xml:"transactionResponse"`
}

type TransactionResponse struct {
	XMLName                     xml.Name              `xml:"transactionResponse"`
	Order                       string                `xml:"orderId"`
	TransactionId               string                `xml:"transactionId"`
	State                       string                `xml:"state"`
	PaymentNetworkResponseCode  string                `xml:"paymentNetworkResponseCode"`
	PaymentNetworkResponseError interface{}           `xml:"paymentNetworkResponseErrorMessage,omitempty"`
	TrazabilityCode             string                `xml:"trazabilityCode"`
	AuthorizationCode           string                `xml:"authorizationCode"`
	PendingReason               interface{}           `xml:"pendingReason,omitempty"`
	ResponseCode                string                `xml:"responseCode"`
	ErrorCode                   interface{}           `xml:"errorCode,omitempty"`
	ResponseMessage             string                `xml:"responseMessage"`
	TransactionDate             interface{}           `xml:"transactionDate,omitempty"`
	TransactionTime             interface{}           `xml:"transactionTime,omitempty"`
	OperationDate               string                `xml:"operationDate"`
	ReferenceQuestionnaire      interface{}           `xml:"referenceQuestionnaire,omitempty"`
	ExtraParameters             ExtraParameters       `xml:"extraParameters"`
	AdditionalInfo              TransactionAdditional `xml:"additionalInfo"`
}

type ExtraParameterss struct {
	XMLName xml.Name               `xml:"extraParameters"`
	Entries []ExtraParametersEntry `xml:"entry"`
}

type ExtraParametersEntry struct {
	Key   string `xml:"string"`
	Value string `xml:"string"`
}

type TransactionAdditional struct {
	XMLName                   xml.Name `xml:"additionalInfo"`
	PaymentNetwork            string   `xml:"paymentNetwork"`
	RejectionType             string   `xml:"rejectionType"`
	ResponseNetworkMessage    string   `xml:"responseNetworkMessage"`
	TravelAgencyAuthorization string   `xml:"travelAgencyAuthorization,omitempty"`
	CardType                  string   `xml:"cardType"`
	TransactionType           string   `xml:"transactionType"`
}
