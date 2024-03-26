package models

type ThreeDomainSecure struct {
	Embedded                     bool   `json:"embedded"`
	ECI                          int    `json:"eci"`
	XID                          string `json:"xid"`
	CAVV                         string `json:"cavv"`
	DirectoryServerTransactionID string `json:"directoryServerTransactionId"`
}
