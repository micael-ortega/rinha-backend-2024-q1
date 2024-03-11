package request

type TransactionRequest struct {
	Value       int    `json:"valor"`
	Description string `json:"descricao"`
	Kind        string `json:"tipo"`
}
