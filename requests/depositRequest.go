package requests

type DepositRequest struct {
	AccountNumber string  `json:"accountNumber"`
	Amount        float64 `json:"amount"`
}
