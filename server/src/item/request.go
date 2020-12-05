package item

import (
	"buycryptos/server/database"
	"encoding/json"
	"os"
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
		q := c.Request.URL.Query()
		t := q["token"][0]
		r, e := database.DB.Query("select id, action_type, title, base_link, description, comment, rating, item_id, user_email, is_approved, created_at from requests " + database.StandardizeQuery(q) + ";")
		defer r.Close()
		code, requests := 200, []*request{}
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
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
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &requests,
		})
	}
}

// RequestPOST export
func RequestPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, code := &request{}, 200
		payload, _ := c.GetRawData()
		json.Unmarshal(payload, &r)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("insert into requests (action_type, title, base_link, description, comment, rating, item_id, user_email) values ($1,$2,$3,$4,$5,$6,$7,$8);", &r.ActionType, &r.Title, &r.BaseLink, &r.Description, &r.Comment, &r.Rating, &r.ItemID, &r.UserEmail)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
		})
	}
}

// RequestPUT export
func RequestPUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		rt, a, t, rID, r, code := q["request_type"][0], q["approved"][0], q["token"][0], q["request_id"][0], &request{}, 200
		payload, _ := c.GetRawData()
		json.Unmarshal(payload, &r)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			// delete request
			if a != "0" {
				switch rt {
				case "create":
					// _, e = tx.Exec("insert into items (value,value) values ($1,$2);", &r.Title, &r.BaseLink)
				case "edit":

				case "delete":
				}
			}
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
		})
	}
}
