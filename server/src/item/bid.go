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
		q, code, bids, em, p := c.Request.URL.Query(), 200, []*bid{}, "", ""
		ID, hasID := q["id"]
		aID, hasaID := q["auction_id"]
		var err interface{}
		if hasID || hasaID {
			if hasID {
				r1, e := database.DB.Query("select email from users where id='" + ID[0] + "' ")
				defer r1.Close()
				if e == nil {
					for r1.Next() {
						r1.Scan(&em)
					}
				} else {
					err, code = string(e.Error()), 500
				}
				p = "where user_email='" + em + "' "
			}
			if hasaID {
				p = "where auction_id='" + aID[0] + "' "
			}
			r2, e := database.DB.Query("select id, auction_id, user_email, is_initial, amount, created_at from bids " + p + database.StandardizeQuery(q) + ";")
			defer r2.Close()
			if e == nil {
				for r2.Next() {
					b := &bid{}
					r2.Scan(&b.ID, &b.AuctionID, &b.UserEmail, &b.IsInitial, &b.Amount, &b.CreatedAt)
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
		} else {
			err, code = "Forbidden", 403
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &bids,
		})
	}
}
