package item

import (
	"buycryptos/server/database"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID          *int      `json:"id"`
	ActionType  *string   `json:"action_type"`
	Title       *string   `json:"title"`
	BaseLink    *string   `json:"base_link"`
	Description *string   `json:"description"`
	Comment     *string   `json:"comment"`
	Rating      *int      `json:"rating"`
	ItemID      *int      `json:"item_id"`
	UserEmail   *int      `json:"user_email"`
	IsApproved  *bool     `json:"is_approved"`
	CreatedAt   time.Time `json:"created_at"`
}

// RequestsGET export
func RequestsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, action_type, title, base_link, description, comment, rating, item_id, user_email, is_approved, created_at from requests " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		requests := []*request{}
		var err interface{}
		if e == nil {
			for r.Next() {
				rq := &request{}
				r.Scan(&rq.ID, &rq.ActionType, &rq.Title, &rq.BaseLink, &rq.Description, &rq.Comment, &rq.Rating, &rq.ItemID, &rq.UserEmail, &rq.IsApproved, &rq.CreatedAt)
				requests = append(requests, rq)
			}
			if len(requests) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Request not found"
				} else {
					err = "No requests"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  requests,
		})
	}
}

// RequestPOST export
func RequestPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, code := &request{}, 200
		payload, _ := c.GetRawData()
		json.Unmarshal(payload, r)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("INSERT INTO requests (action_type, title, base_link, description, comment, rating, item_id, user_email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);", &r.ActionType, &r.Title, &r.BaseLink, &r.Description, &r.Comment, &r.Rating, &r.ItemID, &r.UserEmail)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
		})
	}
}
