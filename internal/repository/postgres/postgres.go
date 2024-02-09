package postgres

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var ErrInvalidDatabaseURL = errors.New("invalid database URL")

type Repository struct {
	db *bun.DB
}

func NewRepository(connection *sql.DB) (*Repository, error) {
	bunDB := bun.NewDB(connection, pgdialect.New())

	return &Repository{db: bunDB}, nil
}

func NewConnection(databaseURL string) (*sql.DB, error) {
	dbConfig, err := pgx.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}
	dbConfig.PreferSimpleProtocol = true
	return stdlib.OpenDB(*dbConfig), nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}
