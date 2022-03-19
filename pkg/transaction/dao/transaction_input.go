package dao

type Transaction_input struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   int64  `json:"amount"`
}

type Transaction_output struct {
	Status       int    `json:"status"`
	Tno          string `json:"transaction_number"`
	Time         string `json:"time"`
	Amount       int64  `json:"amount"`
	SenderName   string `json:"sender"`
	ReceiverName string `json:"receiver"`
}