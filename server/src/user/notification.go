package user

import (
	"buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type notification struct {
	ID        *int      `json:"id"`
	UserID    *int      `json:"user_id"`
	Content   *string   `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsRead    *bool     `json:"is_read"`
}

// NotificationsGET export
func NotificationsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, user_id, content, created_at, is_read from notifications " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		notifications := []*notification{}
		var err interface{}
		if e == nil {
			for r.Next() {
				n := &notification{}
				r.Scan(&n.ID, &n.UserID, &n.Content, &n.CreatedAt, &n.IsRead)
				notifications = append(notifications, n)
			}
			if len(notifications) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Notifications not found"
				} else {
					err = "No notifications"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  notifications,
		})
	}
}
