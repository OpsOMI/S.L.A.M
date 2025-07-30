package txmanagerpkg

import (
	"context"
	"database/sql"
	"errors"
)

type ITXManager interface {
	RunInTx(ctx context.Context, fn func(tx *sql.Tx) error) error
}

type manager struct {
	db *sql.DB
}

func New(db *sql.DB) ITXManager {
	return &manager{db: db}
}

func (tm *manager) RunInTx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}
	return err
}
