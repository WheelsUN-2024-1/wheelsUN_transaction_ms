package models

type Payer struct {
	EmailAddress    string  `json:"emailAddress"`
	MerchantPayerID string  `json:"merchantPayerId"`
	FullName        string  `json:"fullName"`
	BillingAddress  Address `json:"billingAddress"`
	Birthdate       string  `json:"birthdate"`
	ContactPhone    string  `json:"contactPhone"`
	DniNumber       string  `json:"dniNumber"`
	DniType         string  `json:"dniType"`
}
