package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/micael-ortega/crebito/internal/dto"
	"github.com/micael-ortega/crebito/internal/dto/response"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

const (
	FN_CREDIT_TRANSACTION   = "SELECT * FROM fn_credit($1,$2,$3,$4);"
	FN_DEBIT_TRANSACTION    = "SELECT * FROM fn_debit($1,$2,$3,$4);"
	FN_LAST_10_TRANSACTIONS = "SELECT * FROM fn_get_last_transactions($1);"
)

func (r *Repo) CreditTransaction(value int, description string, client_id int) (*dto.TransactionDTO, error) {

	row := r.db.QueryRow(context.Background(), FN_CREDIT_TRANSACTION, client_id, description, "c", value)

	var dto dto.TransactionDTO
	err := row.Scan(
		&dto.Limit,
		&dto.Balance,
		&dto.Code,
	)

	if err != nil {
		return nil, err
	}

	return &dto, nil

}
func (r *Repo) DebitTransaction(value int, description string, client_id int) (*dto.TransactionDTO, error) {

	row := r.db.QueryRow(context.Background(), FN_DEBIT_TRANSACTION, client_id, description, "d", value)

	var dto dto.TransactionDTO
	err := row.Scan(
		&dto.Limit,
		&dto.Balance,
		&dto.Code,
	)

	if err != nil {
		return nil, err
	}

	return &dto, nil

}

func (r *Repo) GetLastTransactions(id int) (*response.BankStatement, error) {
	var res response.BankStatement
	rows, err := r.db.Query(context.Background(), FN_LAST_10_TRANSACTIONS, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	stats := r.db.Stat()

	slog.AnyValue(stats)

	var balance, accountLimit int
	var timestamp time.Time
	var code int

	for rows.Next() {
		var transaction response.Transactions
		err = rows.Scan(
			&balance,
			&accountLimit,
			&timestamp,
			&transaction.Value,
			&transaction.Kind,
			&transaction.Description,
			&transaction.Timestamp,
			&code,
		)
		if code == -1 {
			return nil, fmt.Errorf("user not found")
		}

		if err != nil {
			return nil, err
		}
		if transaction.Value != nil {
			res.LastTransactions = append(res.LastTransactions, transaction)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	res.Balance = response.ClientBalance{
		Balance:   balance,
		Limit:     accountLimit,
		Timestamp: timestamp,
	}

	return &res, nil
}
