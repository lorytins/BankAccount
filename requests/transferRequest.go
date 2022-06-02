package requests

type TransferRequest struct {
	OriginAccountNumber      string  `json:"originAccountNumber"`
	DestinationAccountNumber string  `json:"destinationAccountNumber"`
	Amount                   float64 `json:"amount"`
}
