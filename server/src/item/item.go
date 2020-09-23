package item

import (
	"log"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type item struct {
	ID uuid.UUID `json:"id"`
}

// GET items
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		// test
		queryParams := database.StandardizeQuery(c.Request.URL.Query(), "WHERE")
		itemRows, err := database.DB.Query("SELECT id, ...... FROM items " + queryParams + " ....... ;")
		defer itemRows.Close()
		if err != nil {
			items := []*item{}
			for itemRows.Next() {
				i := &item{}
				itemRows.Scan(&i.ID)
				items = append(items, i)
			}
			c.JSON(200, &gin.H{
				"statusCode": "200",
				"message":    "OK",
				"error":      nil,
				"meta": gin.H{
					"query":       c.Request.URL.Query(),
					"resultCount": len(items),
				},
				"data": items,
			})
		} else {
			c.JSON(500, &gin.H{
				"statusCode": "500",
				"message":    "Internal Server Error",
				"error":      err.Error(),
				"meta": gin.H{
					"query": c.Request.URL.Query(),
				},
			})
			log.Println(err)
		}
	}
}
