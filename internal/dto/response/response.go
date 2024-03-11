package response

import "time"

type TransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type ClientBalance struct {
	Balance   int       `json:"total"`
	Timestamp time.Time `json:"data_extrato"`
	Limit     int       `json:"limite"`
}
type Transactions struct {
	Value       *int       `json:"valor"`
	Kind        *string    `json:"tipo"`
	Description *string    `json:"descricao"`
	Timestamp   *time.Time `json:"realizada_em"`
}

type BankStatement struct {
	Balance          ClientBalance        `json:"saldo"`
	LastTransactions []Transactions `json:"ultimas_transacoes"`
}
