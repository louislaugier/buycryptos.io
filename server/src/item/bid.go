package item

import (
	"buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type bid struct {
	ID        *int      `json:"id"`
	AuctionID *int      `json:"auction_id"`
	UserEmail *int      `json:"user_email"`
	IsInitial *bool     `json:"is_initial"`
	Amount    *int      `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// BidsGET export
func BidsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q, code, bids, ue := c.Request.URL.Query(), 200, []*bid{}, ""
		em, hasEmail := q["user_email"]
		if hasEmail {
			ue = "where user_email='" + em[0] + "' "
		}
		r, e := database.DB.Query("select id, auction_id, user_email, is_initial, amount, created_at from bids " + ue + database.StandardizeQuery(q) + ";")
		defer r.Close()
		var err interface{}
		if e == nil {
			for r.Next() {
				b := &bid{}
				r.Scan(&b.ID, &b.AuctionID, &b.UserEmail, &b.IsInitial, &b.Amount, &b.CreatedAt)
				bids = append(bids, b)
			}
			if len(bids) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Bid not found"
				} else {
					err = "No bids"
				}
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &bids,
		})
	}
}
