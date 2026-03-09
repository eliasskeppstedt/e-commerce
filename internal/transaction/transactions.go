package transaction

import (
	"context"
	"database/sql"
)

type TxManager struct {
	db *sql.DB
}

func NewTxManager(db *sql.DB) *TxManager {
	return &TxManager{db: db}
}

func (tm *TxManager) Begin(ctx context.Context) (*sql.Tx, error) {
	return tm.db.BeginTx(ctx, nil)
}
