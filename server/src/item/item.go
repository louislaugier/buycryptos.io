package item

import (
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID                 int8      `json:"id"`
	Title              string    `json:"title"`
	BaseLink           string    `json:"base_link"`
	RefLink            string    `json:"ref_link"`
	Description        string    `json:"description"`
	IsFeatured         bool      `json:"is_featured"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	RefLinkOwnerUserID int8      `json:"ref_link_owner_user_id"`
	FeaturerUserID     string    `json:"featurer_user_id"`
}

// GET items
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		queryParams := database.StandardizeQuery(c.Request.URL.Query(), "WHERE")
		itemsRows, err := database.DB.Query("SELECT * FROM items " + queryParams + "';")
		defer itemsRows.Close()
		code := 200
		msg := "OK"
		items := []*item{}
		if err != nil {
			for itemsRows.Next() {
				i := &item{}
				itemsRows.Scan(&i.ID)
				items = append(items, i)
			}
			if len(items) == 0 {
				code = 404
				msg = "No items"
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					msg = "Item not found"
				}
			}
		} else {
			code = 500
			msg = "Internal server error"
		}
		c.JSON(code, &gin.H{
			"statusCode": code,
			"message":    msg,
			"error":      err.Error(),
			"meta": gin.H{
				"query": c.Request.URL.Query(),
			},
			"data": items,
		})
	}
}
