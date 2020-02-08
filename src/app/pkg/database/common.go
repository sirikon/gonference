package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gonference/pkg/utils"
	"strconv"
	"strings"
)

func query(conn *pgxpool.Pool, sql string, args ...interface{}) pgx.Rows {
	rows, err := conn.Query(context.Background(), sql, args...); utils.Check(err)
	return rows
}

func exec(conn *pgxpool.Pool, sql string, args ...interface{})  {
	fmt.Println(sql)
	fmt.Println(args)
	_, err := conn.Exec(context.Background(), sql, args...); utils.Check(err)
}

func selectQuery(conn *pgxpool.Pool, fields string, table string, extra string, args ...interface{}) pgx.Rows {
	return query(conn, "SELECT " + fields + " FROM " + table + " " + extra, args...)
}

func insertQuery(conn *pgxpool.Pool, fields string, table string, args ...interface{}) {
	valuesFragment := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		valuesFragment[i] = "$" + strconv.Itoa(i + 1)
	}

	exec(conn,
		"INSERT INTO " + table + " (" + fields + ") VALUES (" + strings.Join(valuesFragment, ", ") + ")",
		args...)
}
