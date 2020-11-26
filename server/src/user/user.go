package user

import (
	"encoding/json"
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type user struct {
	ID          *int      `json:"id"`
	Email       *string   `json:"email"`
	Password    *string   `json:"password"`
	Balance     *int      `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	IsActivated *bool     `json:"is_activated"`
	LastIP      string    `json:"last_ip"`
	IsAdmin     *bool     `json:"is_admin"`
	IsDeleted   *bool     `json:"is_deleted"`
	Token       uuid.UUID `json:"token"`
}

// GET users
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		r, e := database.DB.Query("select id, email, password, balance, created_at, is_activated, last_ip, is_admin, is_deleted from users " + database.StandardizeQuery(q) + ";")
		defer r.Close()
		code, users := 200, []*user{}
		var err interface{}
		if e == nil {
			for r.Next() {
				i := &user{}
				r.Scan(&i.ID, &i.Email, &i.Password, &i.Balance, &i.CreatedAt, &i.IsActivated, &i.LastIP, &i.IsAdmin, &i.IsDeleted)
				users = append(users, i)
			}
			if len(users) == 0 {
				code = 404
				_, hasID := q["id"]
				if hasID {
					err = "User not found"
				} else {
					err = "No users"
				}
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  users,
		})
	}
}

// POST user
func POST() func(c *gin.Context) {
	return func(c *gin.Context) {
		u, code := &user{}, 200
		payload, _ := c.GetRawData()
		u.Token, u.LastIP = uuid.New(), c.ClientIP()
		json.Unmarshal(payload, u)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("INSERT INTO users (email, password, last_ip, token) VALUES ($1, $2, $3, $4);", &u.Email, &u.Password, &u.LastIP, &u.Token)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
		})
	}
}
