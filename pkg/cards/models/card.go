package models

type Account struct {
	MONEY_ACCOUNT_BALANCE float64 `json:"balance"`
	MONEY_ACCOUNT_ID      string  `json:"card_number"`
	SECURITY_CODE         string  `json:"security_code"`
	IS_ACTIVE             bool    `json:"enabled"`
	ExpireOn              string  `json:"expire_on"`
	Type                  string  `json:"type"`
	CardHolder            string  `json:"card_holder"`
}
