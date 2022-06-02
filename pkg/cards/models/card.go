package models

type Card struct {
	Type         string `json:"type"`
	CardNumber   string `json:"card_number"`
	ExpireOn     string `json:"expire_on"`
	SecurityCode string `json:"security_code"`
	CardHolder   string `json:"card_holder"`
	Enabled      bool   `json:"enabled"`
}
