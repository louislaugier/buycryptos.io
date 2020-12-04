package item

import (
	"buycryptos/server/database"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type rating struct {
	ID        *int    `json:"id"`
	ItemID    *int    `json:"item_id"`
	UserEmail *int    `json:"user_email"`
	Rating    *int8   `json:"rating"`
	Comment   *string `json:"comment"`
}

// RatingsGET export
func RatingsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		i := ""
		itemID, hasItemID := c.Request.URL.Query()["item_id"]
		if hasItemID {
			i = "where item_id='" + itemID[0] + "' "
		}
		r, e := database.DB.Query("select id, item_id, user_email, rating, comment from ratings " + i + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code, ratings := 200, []*rating{}
		var err interface{}
		if e == nil {
			for r.Next() {
				rtg := &rating{}
				r.Scan(&rtg.ID, &rtg.ItemID, &rtg.UserEmail, &rtg.Rating, &rtg.Comment)
				ratings = append(ratings, rtg)
			}
			if len(ratings) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Rating not found"
				} else {
					err = "No ratings"
				}
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &ratings,
		})
	}
}

// AverageRatingGET export
func AverageRatingGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select rating from ratings where item_id='" + c.Request.URL.Query()["item_id"][0] + "';")
		defer r.Close()
		code, ratings, t := 200, []*float64{}, 0.00
		var err, a interface{}
		if e == nil {
			for r.Next() {
				rtg := 0.00
				r.Scan(&rtg)
				t, ratings = t+rtg, append(ratings, &rtg)
			}
			a = t / float64(len(ratings))
			if len(ratings) == 0 {
				code, a, err = 404, nil, "No ratings"
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &a,
		})
	}
}

// RatingPOST export
func RatingPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, r, code := q["id"][0], &rating{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &r)
		// get user_email by id
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("insert into ratings (item_id, user_email, rating, comment) values ($1,$2,$3,$4);", &r.ItemID, em, &r.Rating, &r.Comment)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &r,
		})
	}
}

// RatingPUT export
func RatingPUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, rID, r, code := q["id"][0], q["rating_id"][0], &rating{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &r)
		// get user_email by id
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("update ratings set item_id=$1, user_email=$2, rating=$3, comment=$4 where id='"+rID+"';", &r.ItemID, em, &r.Rating, &r.Comment)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &r,
		})
	}
}

// RatingDELETE export
func RatingDELETE() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, rID := q["id"][0], q["rating_id"][0]
		i, code := &item{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &i)
		// get user_email by id
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil {
			_, e = tx.Exec("delete from ratings where id='" + rID + "';")
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &i,
		})
	}
}
