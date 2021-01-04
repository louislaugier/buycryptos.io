package item

import (
	"github.com/louislaugier/buycryptos/server/database"
	"encoding/json"
	"os"
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
		q, p, em, code, featurings := c.Request.URL.Query(), "", "", 200, []*featuring{}
		t, hasToken := q["token"]
		ID, hasID := q["id"]
		var err interface{}
		if (hasToken && t[0] == os.Getenv("ADMIN_TOKEN")) || hasID {
			if hasID {
				r1, e := database.DB.Query("select  email from users where id='" + ID[0] + "'" + database.StandardizeQuery(q) + ";")
				defer r1.Close()
				if e == nil {
					for r1.Next() {
						r1.Scan(&em)
					}
					p = "where user_email='" + em + "' "
				} else {
					err, code = string(e.Error()), 500
				}
			}
			r2, e := database.DB.Query("select id, user_email, item_id, created_at from featurings " + p + database.StandardizeQuery(q) + ";")
			defer r2.Close()
			if e == nil {
				for r2.Next() {
					f := &featuring{}
					r2.Scan(&f.ID, &f.UserEmail, &f.ItemID, &f.CreatedAt)
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
		} else {
			err, code = "Forbidden", 403
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &featurings,
		})
	}
}

// FeaturingPUT export
func FeaturingPUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, t, a := q["id"][0], q["token"][0], q["action"][0]
		i, code := &item{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &i)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			if a == "unfeature" {
				_, e = tx.Exec("update items set featurer_user_email='' where id='" + ID + "';")
			} else {
				_, e = tx.Exec("update items set featurer_user_email='" + os.Getenv("ADMIN_EMAIL") + "' where id='" + ID + "';")
			}
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
