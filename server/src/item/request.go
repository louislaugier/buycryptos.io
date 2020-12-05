package item

import (
	"buycryptos/server/database"
	"encoding/json"
	"os"
	"strconv"
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
	CategoryID  *int      `json:"category_id"`
	ImagePath   *string   `json:"image_path"`
	UserEmail   *int      `json:"user_email"`
	IsApproved  *bool     `json:"is_approved"`
	CreatedAt   time.Time `json:"created_at"`
}

// RequestsGET export
func RequestsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		t := q["token"][0]
		r, e := database.DB.Query("select id, action_type, title, base_link, description, comment, rating, item_id, category_id, image_path, user_email, is_approved, created_at from requests " + database.StandardizeQuery(q) + ";")
		defer r.Close()
		code, requests := 200, []*request{}
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			for r.Next() {
				rq := &request{}
				r.Scan(&rq.ID, &rq.ActionType, &rq.Title, &rq.BaseLink, &rq.Description, &rq.Comment, &rq.Rating, &rq.ItemID, &rq.CategoryID, &rq.ImagePath, &rq.UserEmail, &rq.IsApproved, &rq.CreatedAt)
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
			_, e = tx.Exec("insert into requests (action_type, title, base_link, description, comment, rating, item_id, category_id, image_path, user_email) values ($1,$2,$3,$4,$5,$6,$7,$8,$9);", &r.ActionType, &r.Title, &r.BaseLink, &r.Description, &r.Comment, &r.Rating, &r.ItemID, &r.CategoryID, &r.ImagePath, &r.UserEmail)
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
		rt, a, t, r, code := q["request_type"][0], q["approve"][0], q["token"][0], &request{}, 200
		payload, _ := c.GetRawData()
		json.Unmarshal(payload, &r)
		tx1, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx1.Exec("delete from requests where id='" + strconv.Itoa(*r.ID) + "';")
			tx1.Commit()
			if a != "0" {
				tx2, _ := database.DB.Begin()
				switch rt {
				case "create":
					_, e = tx2.Exec("insert into items (title, base_link, description, category_id, image_path) values ($1,$2,$3,$4,$5);", &r.Title, &r.BaseLink, &r.Description, &r.CategoryID, &r.ImagePath)
				case "edit":
					_, e = tx2.Exec("update items set title=$1, base_link=$2, description=$3, updated_at=$4,category_id=$5, image_path=$6 where id='"+strconv.Itoa(*r.ID)+"';", &r.Title, &r.BaseLink, &r.Description, time.Now(), &r.CategoryID, &r.ImagePath)
				case "delete":
					_, e = tx2.Exec("delete from items where id='" + strconv.Itoa(*r.ID) + "';")
				}
				tx2.Commit()
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
		})
	}
}
