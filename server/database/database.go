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
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// StandardizeQuery from HTTP to SQL
func StandardizeQuery(query url.Values, operator string) string {
	ID := ""
	orderBy := ""
	order := ""
	limit := ""
	offset := ""
	if _, i := query["id"]; i {
		orderBy = operator + " id='" + query["id"][0] + "' "
	}
	if _, i := query["orderby"]; i {
		orderBy = "ORDER BY " + query["orderby"][0] + " "
	}
	if _, i := query["order"]; i {
		order = query["order"][0] + " "
	}
	if _, i := query["limit"]; i {
		limit = "LIMIT " + query["limit"][0] + " "
	}
	if _, i := query["offset"]; i {
		offset = "OFFSET " + query["offset"][0] + " "
	}
	return strings.TrimSuffix(" "+ID+orderBy+order+limit+offset, " ")
}
