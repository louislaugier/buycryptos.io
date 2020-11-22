package item

import (
	"buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type featuring struct {
	ID        *int      `json:"id"`
	UserID    *int      `json:"user_id"`
	ItemID    *int      `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
}

// FeaturingsGET export
func FeaturingsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, user_id, item_id, created_at from featurings " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		featurings := []*featuring{}
		var err interface{}
		if e == nil {
			for r.Next() {
				f := &featuring{}
				r.Scan(&f.ID, &f.UserID, &f.ItemID, &f.CreatedAt)
				featurings = append(featurings, f)
			}
			if len(featurings) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Featuring not found"
				} else {
					err = "No featurings"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  featurings,
		})
	}
}
