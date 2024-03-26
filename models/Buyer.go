package models

type Buyer struct {
    MerchantBuyerID string  `json:"merchantBuyerId"`
    FullName        string  `json:"fullName"`
    EmailAddress    string  `json:"emailAddress"`
    DniNumber       string  `json:"dniNumber"`
    ShippingAddress Address `json:"shippingAddress"`
}
