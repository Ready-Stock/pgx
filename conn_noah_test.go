package pgx_test

import (
	"testing"
	"github.com/readystock/pgx"
)

var (
	DefaultConnConfig = pgx.ConnConfig{
		Host:"127.0.0.1",
		Port: 5432,
		User: "postgres",
		Password:"Spring!2016",
	}
)

func Test_TestQueryEX(t *testing.T) {
	conn, _ := pgx.Connect(DefaultConnConfig)
	conn.Query("SELECT 1")

}