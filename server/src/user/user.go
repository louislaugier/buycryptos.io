package user

import (
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID          int8      `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Balance     int8      `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	IsActivated bool      `json:"is_activated"`
	LastIP      string    `json:"last_ip"`
	IsAdmin     bool      `json:"is_admin"`
	IsDeleted   bool      `json:"is_deleted"`
}

// GET users
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		queryParams := database.StandardizeQuery(c.Request.URL.Query(), "WHERE")
		userRows, err := database.DB.Query("SELECT * FROM users " + queryParams + "';")
		defer userRows.Close()
		code := 200
		msg := "OK"
		users := []*user{}
		if err != nil {
			for userRows.Next() {
				u := &user{}
				userRows.Scan(&u.ID)
				users = append(users, u)
			}
			if len(users) == 0 {
				code = 404
				msg = "No users"
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					msg = "User not found"
				}
			}
		} else {
			code = 500
			msg = "Internal server error"
		}
		c.JSON(code, &gin.H{
			"statusCode": code,
			"message":    msg,
			"error":      err.Error(),
			"meta": gin.H{
				"query": c.Request.URL.Query(),
			},
			"data": users,
		})
	}
}
