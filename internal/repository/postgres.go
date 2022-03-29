package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type goCommerceRepo struct {
	DB *pgxpool.Pool
}

func NewGoCommerceRepo(connString string) (*goCommerceRepo, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return &goCommerceRepo{DB: db}, nil
}
