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
		r, e := database.DB.Query("select id, item_id, user_id, rating, comment from ratings " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		ratings := []*rating{}
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
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  ratings,
		})
	}
}
