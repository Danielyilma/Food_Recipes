package models

type Payload struct {
	Amount      string `json:"amount"`
	PhoneNumber string `json:"phone_number"`
	ReturnUrl   string `json:"return_url"`
}
