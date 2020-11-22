package database

import (
	"database/sql"
	"net/url"
	"strings"

	_ "github.com/lib/pq" // Postgres driver
)

// DB is a SQL database pointer
var DB = connect()

func connect() *sql.DB {
	// if err := godotenv.Load(); err != nil {
	// 	panic(err)
	// }
	db, _ := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/local_dev?sslmode=disable")
	e := db.Ping()
	if e != nil {
		panic(e)
	}
	return db
}

// StandardizeQuery from URL to SQL
func StandardizeQuery(q url.Values) string {
	ID := ""
	orderBy := ""
	order := ""
	limit := ""
	offset := ""
	if _, i := q["id"]; i {
		orderBy = "where id='" + q["id"][0] + "' "
	}
	if _, i := q["orderby"]; i {
		orderBy = "order by " + q["orderby"][0] + " "
	}
	if _, i := q["order"]; i {
		order = q["order"][0] + " "
	}
	if _, i := q["limit"]; i {
		limit = "limit " + q["limit"][0] + " "
	}
	if _, i := q["offset"]; i {
		offset = "offset " + q["offset"][0] + " "
	}
	return strings.TrimSuffix(ID+orderBy+order+limit+offset, " ")
}
