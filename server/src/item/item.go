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
	CategoryID         *int      `json:"category_id"`
}

// GET items
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q, cID := c.Request.URL.Query(), ""
		if _, i := q["category_id"]; i {
			cID = "where category_id='" + q["category_id"][0] + "' "
		}
		r, e := database.DB.Query("select id, title, base_link, ref_link, description, is_featured, created_at, updated_at, ref_link_owner_user_id, featurer_user_id from items " + cID + database.StandardizeQuery(q) + ";")
		defer r.Close()
		code, items := 200, []*item{}
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
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  items,
		})
	}
}
