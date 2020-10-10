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

// GET user by ID
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		queryParams := database.StandardizeQuery(c.Request.URL.Query(), "WHERE")
		userRow, err := database.DB.Query("SELECT * FROM users " + queryParams + "';")
		defer userRow.Close()
		code := 200
		msg := "OK"
		u := &user{}
		if err != nil {
			for userRow.Next() {
				userRow.Scan(&u.ID)
			}
			if u == nil {
				code = 404
				msg = "User not found"
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
			"data": u,
		})
	}
}
