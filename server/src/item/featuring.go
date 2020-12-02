package item

import (
	"buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type featuring struct {
	ID        *int      `json:"id"`
	UserEmail *int      `json:"user_email"`
	ItemID    *int      `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
}

// FeaturingsGET export
func FeaturingsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q, code, featurings, ue := c.Request.URL.Query(), 200, []*featuring{}, ""
		em, hasEmail := q["user_email"]
		if hasEmail {
			ue = "where user_email='" + em[0] + "' "
		}
		r, e := database.DB.Query("select id, user_email, item_id, created_at from featurings " + ue + database.StandardizeQuery(q) + ";")
		defer r.Close()
		var err interface{}
		if e == nil {
			for r.Next() {
				f := &featuring{}
				r.Scan(&f.ID, &f.UserEmail, &f.ItemID, &f.CreatedAt)
				featurings = append(featurings, f)
			}
			if len(featurings) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Featuring not found"
				} else {
					err = "No featuring history"
				}
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &featurings,
		})
	}
}
