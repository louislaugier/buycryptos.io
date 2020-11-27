package item

import (
	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type rating struct {
	ID      *int    `json:"id"`
	ItemID  *int    `json:"item_id"`
	UserID  *int    `json:"user_id"`
	Rating  *int8   `json:"rating"`
	Comment *string `json:"comment"`
}

// RatingsGET export
func RatingsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		i := ""
		itemID, hasItemID := c.Request.URL.Query()["item_id"]
		if hasItemID {
			i = "where item_id='" + itemID[0] + "' "
		}
		r, e := database.DB.Query("select id, item_id, user_id, rating, comment from ratings " + i + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code, ratings := 200, []*rating{}
		var err interface{}
		if e == nil {
			for r.Next() {
				rtg := &rating{}
				r.Scan(&rtg.ID, &rtg.ItemID, &rtg.UserID, &rtg.Rating, &rtg.Comment)
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
			"error": err,
			"data":  ratings,
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
			"error": err,
			"data":  a,
		})
	}
}
