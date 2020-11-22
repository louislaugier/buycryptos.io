package user

import (
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID          *int      `json:"id"`
	Email       *string   `json:"email"`
	Password    *string   `json:"password"`
	Balance     *int      `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	IsActivated *bool     `json:"is_activated"`
	LastIP      *string   `json:"last_ip"`
	IsAdmin     *bool     `json:"is_admin"`
	IsDeleted   *bool     `json:"is_deleted"`
}

// GET users
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := database.StandardizeQuery(c.Request.URL.Query())
		r, e := database.DB.Query("select id, email, password, balance, created_at, is_activated, last_ip, is_admin, is_deleted from users " + q + ";")
		defer r.Close()
		code := 200
		users := []*user{}
		var err interface{}
		if e == nil {
			for r.Next() {
				i := &user{}
				r.Scan(&i.ID, &i.Email, &i.Password, &i.Balance, &i.CreatedAt, &i.IsActivated, &i.LastIP, &i.IsAdmin, &i.IsDeleted)
				users = append(users, i)
			}
			if len(users) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "User not found"
				} else {
					err = "No users"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  users,
		})
	}
}
