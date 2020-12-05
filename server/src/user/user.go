package user

import (
	"encoding/json"
	"os"
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type user struct {
	ID              uuid.UUID `json:"id"`
	Email           *string   `json:"email"`
	Password        *string   `json:"password"`
	Balance         *int      `json:"balance"`
	CreatedAt       time.Time `json:"created_at"`
	IsEmailVerified bool      `json:"is_email_verified"`
	LastIP          string    `json:"last_ip"`
	IsAdmin         *bool     `json:"is_admin"`
}

// GET export
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		q := c.Request.URL.Query()
		ID, code, users, em := q["token"][0], 200, []*user{}, ""
		var err interface{}
		if ID == os.Getenv("ADMIN_TOKEN") {
			email, hasEmail := q["email"]
			if hasEmail {
				em = "where email='" + email[0] + "' "
			}
			r, e := database.DB.Query("select id, email, balance, created_at, is_email_verified, last_ip, is_admin from users " + em + database.StandardizeQuery(q) + ";")
			defer r.Close()
			if e == nil {
				for r.Next() {
					u := &user{}
					r.Scan(&u.ID, &u.Email, &u.Balance, &u.CreatedAt, &u.IsEmailVerified, &u.LastIP, &u.IsAdmin)
					users = append(users, u)
				}
				if len(users) == 0 {
					code = 404
					if hasEmail {
						err = "User not found"
					} else {
						err = "No users"
					}
				}
			} else {
				err, code = string(e.Error()), 500
			}
		} else {
			err, code = "Forbidden", 403
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &users,
		})
	}
}

// PUT export
func PUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		q := c.Request.URL.Query()
		ID, code, u := q["id"][0], 200, &user{}
		p, _ := c.GetRawData()
		json.Unmarshal(p, u)
		var err interface{}
		if ID == os.Getenv("ADMIN_TOKEN") {
			tx, e := database.DB.Begin()
			if e == nil {
				if *u.IsAdmin {
					_, e = tx.Exec("update users set id=$1 email=$2, balance=$3, is_email_verified=$4, is_admin=$5) values ($1,$2,$3,$4,$5);", os.Getenv("ADMIN_TOKEN"), &u.Email, &u.Balance, &u.IsEmailVerified, &u.IsAdmin)
				} else {
					_, e = tx.Exec("update users set email=$1, balance=$2, is_email_verified=$3, is_admin=$4) values ($1,$2,$3,$4);", &u.Email, &u.Balance, &u.IsEmailVerified, &u.IsAdmin)
				}
				tx.Commit()
			} else {
				err, code = string(e.Error()), 500
			}
		} else {
			err, code = "Forbidden", 403
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &u,
		})
	}
}

// Login export
func Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		e, p, code, u := q["email"][0], q["password"][0], 200, &user{}
		var err interface{}
		r, err := database.DB.Query("select id, is_email_verified from users where email='" + e + "' and password='" + p + "';")
		defer r.Close()
		if err == nil {
			for r.Next() {
				r.Scan(&u.ID, &u.IsEmailVerified)
			}
		} else {
			code = 500
		}
		if !u.IsEmailVerified {
			err, code = "Account is not verified", 403
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &u.ID,
		})
	}
}

// ProfileGET export
func ProfileGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, code, u := q["id"][0], 200, &user{}
		var err interface{}
		r, e := database.DB.Query("select email, password, balance, created_at from users where id='" + ID + "';")
		defer r.Close()
		if e == nil {
			for r.Next() {
				r.Scan(&u.Email, &u.Password, &u.Balance, &u.CreatedAt)
			}
			if *u == (user{}) {
				code = 404
				err = "Profile not found"
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &u,
		})
	}
}

// PasswordUpdate export
func PasswordUpdate() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, code, p := q["id"][0], 200, q["password"][0]
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("update users set password=$1 where id='"+ID+"';", &p)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &p,
		})
	}
}

// POST export
func POST() func(c *gin.Context) {
	return func(c *gin.Context) {
		u, code := &user{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &u)
		u.ID, u.LastIP = uuid.New(), c.ClientIP()
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("insert into users (id, email, password, last_ip) values ($1, $2, $3, $4);", &u.ID, &u.Email, &u.Password, &u.LastIP)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &u,
		})
	}
}

// DELETE export
func DELETE() func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		q := c.Request.URL.Query()
		ID, em, code, u := q["id"][0], q["email"][0], 200, &user{}
		var err interface{}
		r, e := database.DB.Query("select id, email from users where email='" + em + "';")
		defer r.Close()
		if e == nil {
			for r.Next() {
				r.Scan(&u.ID, &u.Email)
			}
		} else {
			err, code = string(e.Error()), 500
		}
		if *u.Email != "" {
			if ID == os.Getenv("ADMIN_TOKEN") || ID == u.ID.String() {
				tx, e := database.DB.Begin()
				if e == nil {
					_, e = tx.Exec("delete from users where email='" + em + "';")
					tx.Commit()
				} else {
					err, code = string(e.Error()), 500
				}
			} else {
				err, code = "Forbidden", 403
			}
		} else {
			err, code = "No such user", 404
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &u,
		})
	}
}
