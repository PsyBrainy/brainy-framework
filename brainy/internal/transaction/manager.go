package transaction

import (
	"context"
	"database/sql"
)

type TransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) Begin(ctx context.Context) (*sql.Tx, error) {
	return tm.db.BeginTx(ctx, nil)
}

func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := tm.Begin(ctx)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
