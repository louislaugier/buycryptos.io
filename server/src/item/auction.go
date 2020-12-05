package item

import (
	"buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type auction struct {
	ID            *int      `json:"id"`
	ItemID        *int      `json:"item_id"`
	IncrementRate time.Time `json:"increment_rate"`
	WinnerEmail   time.Time `json:"winner_email"`
}

// AuctionsGET export
func AuctionsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, item_id, increment_rate, winner_email from auctions " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code, auctions := 200, []*auction{}
		var err interface{}
		if e == nil {
			for r.Next() {
				a := &auction{}
				r.Scan(&a.ID, &a.ItemID, &a.IncrementRate, &a.WinnerEmail)
				auctions = append(auctions, a)
			}
			if len(auctions) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Auction not found"
				} else {
					err = "No auctions"
				}
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &auctions,
		})
	}
}
