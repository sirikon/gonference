package client

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/utils"
	"strconv"
	"strings"
)

func GetDBClient(connectionString string) *DBClient {
	return &DBClient{connPool:getPGXPool(connectionString)}
}

type DBClient struct {
	connPool *pgxpool.Pool
}

func (dbc *DBClient) Exec(sql string, args ...interface{}) {
	_, err := dbc.connPool.Exec(context.Background(), sql, args...); utils.Check(err)
}

func (dbc *DBClient) Query(sql string, args ...interface{}) pgx.Rows {
	rows, err := dbc.connPool.Query(context.Background(), sql, args...); utils.Check(err)
	return rows
}

func (dbc *DBClient) Select(fields string, table string, extra string, args ...interface{}) pgx.Rows {
	return dbc.Query("SELECT " + fields + " FROM " + table + " " + extra, args...)
}

func (dbc *DBClient) Insert(fields string, table string, args ...interface{}) {
	valuesPlaceholders := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		valuesPlaceholders[i] = "$" + strconv.Itoa(i + 1)
	}
	valuesFragment := strings.Join(valuesPlaceholders, ", ")
	dbc.Exec("INSERT INTO " + table + " (" + fields + ") VALUES (" + valuesFragment + ")", args...)
}
