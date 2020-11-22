package item

import (
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID                 *int      `json:"id"`
	Title              *string   `json:"title"`
	BaseLink           *string   `json:"base_link"`
	RefLink            *string   `json:"ref_link"`
	Description        *string   `json:"description"`
	IsFeatured         *bool     `json:"is_featured"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	RefLinkOwnerUserID *int      `json:"ref_link_owner_user_id"`
	FeaturerUserID     *string   `json:"featurer_user_id"`
}

// GET items
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, title, base_link, ref_link, description, is_featured, created_at, updated_at, ref_link_owner_user_id, featurer_user_id from items " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		items := []*item{}
		var err interface{}
		if e == nil {
			for r.Next() {
				i := &item{}
				r.Scan(&i.ID, &i.Title, &i.BaseLink, &i.RefLink, &i.Description, &i.IsFeatured, &i.CreatedAt, &i.UpdatedAt, &i.RefLinkOwnerUserID, &i.FeaturerUserID)
				items = append(items, i)
			}
			if len(items) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Item not found"
				} else {
					err = "No items"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  items,
		})
	}
}
