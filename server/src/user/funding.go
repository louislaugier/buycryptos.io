package user

import (
	"github.com/louislaugier/buycryptos/server/database"
	"time"

	"github.com/gin-gonic/gin"
)

type funding struct {
	ID            *int      `json:"id"`
	UserEmail     *int      `json:"user_email"`
	Amount        *int      `json:"amount"`
	PaymentMethod *string   `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
}

// FundingsGET export
func FundingsGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q, code, fundings, ue := c.Request.URL.Query(), 200, []*funding{}, ""
		em, hasEmail := q["user_email"]
		if hasEmail {
			ue = "where user_email='" + em[0] + "' "
		}
		r, e := database.DB.Query("select id, action_type, title, base_link, description, comment, rating, item_id, user_email, is_approved, created_at from fundings " + ue + database.StandardizeQuery(q) + ";")
		defer r.Close()
		var err interface{}
		if e == nil {
			for r.Next() {
				f := &funding{}
				r.Scan(&f.ID, &f.UserEmail, &f.Amount, &f.PaymentMethod, &f.CreatedAt)
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
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &fundings,
		})
	}
}
