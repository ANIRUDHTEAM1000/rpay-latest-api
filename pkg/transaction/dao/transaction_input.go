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

type Transaction struct {
	Name              string  `json:"name"`
	Amount            float64 `json:"amount"`
	Date              string  `json:"date"`
	TransactionNumber string  `json:"transaction_number"`
}

type TransactionsList struct {
	Status       int           `json:"status"`
	Transactions []Transaction `json:"transactions"`
	TotalPages   float64       `json:"total_pages"`
}
