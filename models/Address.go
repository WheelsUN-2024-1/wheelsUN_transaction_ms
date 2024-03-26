package models

type Address struct {
    Street1    string `json:"street1"`
    Street2    string `json:"street2"`
    City       string `json:"city"`
    State      string `json:"state"`
    Country    string `json:"country"`
    PostalCode string `json:"postalCode"`
    Phone      string `json:"phone"`
}
