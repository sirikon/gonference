package client

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/utils"
)

func getPGXPool(connectionString string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), connectionString); utils.Check(err)
	return pool
}
