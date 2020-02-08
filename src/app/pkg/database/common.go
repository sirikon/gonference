package database

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/utils"
	"strconv"
	"strings"
)

func selectQuery(conn *pgxpool.Pool, fields string, table string, extra string, args ...interface{}) pgx.Rows {
	sql := "SELECT " + fields + " FROM " + table + " " + extra
	rows, err := conn.Query(context.Background(), sql, args...); utils.Check(err)
	return rows
}

func insertQuery(conn *pgxpool.Pool, fields string, table string, args ...interface{}) {
	valuesFragment := make([]string, len(args))
	for i := range args {
		valuesFragment[i] = "$" + strconv.Itoa(i + 1)
	}
	sql := "INSERT INTO " + table + " (" + fields + ") VALUES (" + strings.Join(valuesFragment, ", ") + ")"
	_, err := conn.Exec(context.Background(), sql, args...); utils.Check(err)
}
