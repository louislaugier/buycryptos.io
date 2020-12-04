package user

import (
	"buycryptos/server/database"
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type notification struct {
	ID        *int      `json:"id"`
	UserEmail *string   `json:"user_email"`
	Content   *string   `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	IsRead    *bool     `json:"is_read"`
}

// NotificationsGET export
func NotificationsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		em := q["user_email"][0]
		r, e := database.DB.Query("select id, user_email, content, created_at, is_read from notifications where user_email='" + em + "'" + database.StandardizeQuery(q) + ";")
		defer r.Close()
		code, notifications := 200, []*notification{}
		var err interface{}
		if e == nil {
			for r.Next() {
				n := &notification{}
				r.Scan(&n.ID, &n.UserEmail, &n.Content, &n.CreatedAt, &n.IsRead)
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
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  notifications,
		})
	}
}

// NotificationPOST export
func NotificationPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		t, n, code := q["token"][0], &notification{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &n)
		tx, e := database.DB.Begin()
		var err interface{}
		if t == os.Getenv("ADMIN_TOKEN") {
			if e == nil {
				_, e = tx.Exec("insert into notifications (user_email, content) values ($1,$2);", &n.UserEmail, &n.Content)
				tx.Commit()
			} else {
				err, code = string(e.Error()), 500
			}
		} else {
			err, code = "Forbidden", 403
		}

		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &n,
		})
	}
}

// NotificationPUT export
func NotificationPUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, em, nID, em2, code := q["id"][0], "", q["notification_id"][0], "", 200
		r, e := database.DB.Query("select user_email from users where id='" + ID + "';")
		defer r.Close()
		var err interface{}
		if e == nil {
			for r.Next() {
				r.Scan(&em2)
			}
		} else {
			err, code = string(e.Error()), 500
		}
		tx, e := database.DB.Begin()
		if em == em2 {
			if e == nil {
				_, e = tx.Exec("update notifications set is_read=true where id='" + nID + "';")
				tx.Commit()
			} else {
				err, code = string(e.Error()), 500
			}
		}

		c.JSON(code, &gin.H{
			"error": &err,
			"data":  true,
		})
	}
}
