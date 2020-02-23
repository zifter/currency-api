package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zifter/currency/service-logic/internal/types"
)

const (
	selectQuery = `SELECT id FROM rate_post WHERE postdate = $1 and msg = $2;`
	insertQuery = `INSERT INTO rate_post (postDate, msg) values($1, $2);`
)

type DBRepo struct {
	db *sqlx.DB
}

func NewDBRepo(db *sqlx.DB) *DBRepo {
	return &DBRepo{db}
}

func (r *DBRepo) IsEntryExists(ctx context.Context, e *types.RatePost) (bool, error) {
	found := []types.RatePost{}
	err := r.db.SelectContext(ctx, &found, selectQuery, e.Date, e.Msg)
	if err != nil {
		return false, fmt.Errorf("cant check existance: %w", err)
	}

	return len(found) != 0, nil
}

func (r *DBRepo) CreateEntry(ctx context.Context, e *types.RatePost) error {
	_, err := r.db.ExecContext(ctx, insertQuery, e.Date, e.Msg)
	if err != nil {
		return fmt.Errorf("cant insert %v: %w", e, err)
	}

	return nil
}
