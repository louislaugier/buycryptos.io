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
	ID, orderBy, order, limit, offset := "", "", "", "", ""
	if _, i := q["id"]; i {
		ID = "where id='" + q["id"][0] + "' "
	}
	// if _, i := q["user_id"]; i {
	// 	ID = "where user_id='" + q["user_id"][0] + "' "
	// }
	// if _, i := q["item_id"]; i {
	// 	ID = "where item_id='" + q["item_id"][0] + "' "
	// }
	// if _, i := q["auction_id"]; i {
	// 	ID = "where auction_id='" + q["auction_id"][0] + "' "
	// }
	// if _, i := q["winner_user_id"]; i {
	// 	ID = "where winner_user_id='" + q["winner_user_id"][0] + "' "
	// }
	// if _, i := q["ref_link_owner_user_id"]; i {
	// 	ID = "where ref_link_owner_user_id='" + q["ref_link_owner_user_id"][0] + "' "
	// }
	// if _, i := q["featurer_user_id"]; i {
	// 	ID = "where featurer_user_id='" + q["featurer_user_id"][0] + "' "
	// }
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
