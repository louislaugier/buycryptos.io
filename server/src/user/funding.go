package user

import (
	"buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type funding struct {
	ID            *int      `json:"id"`
	UserID        *int      `json:"user_id"`
	Amount        *int      `json:"amount"`
	PaymentMethod *string   `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
}

// FundingsGET export
func FundingsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, action_type, title, base_link, description, comment, rating, item_id, user_id, is_approved, created_at from fundings " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		fundings := []*funding{}
		var err interface{}
		if e == nil {
			for r.Next() {
				f := &funding{}
				r.Scan(&f.ID, &f.UserID, &f.Amount, &f.PaymentMethod, &f.CreatedAt)
				fundings = append(fundings, f)
			}
			if len(fundings) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Funding not found"
				} else {
					err = "No fundings"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  fundings,
		})
	}
}
