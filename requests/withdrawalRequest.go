package requests

type WithdrawalRequest struct {
	AccountNumber string  `json:"accountNumber"`
	Amount        float64 `json:"amount"`
}
