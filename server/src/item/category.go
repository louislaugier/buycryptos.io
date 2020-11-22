package item

import (
	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type category struct {
	ID          *int    `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// CategoriesGET export
func CategoriesGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		r, e := database.DB.Query("select id, title, description from categories " + database.StandardizeQuery(c.Request.URL.Query()) + ";")
		defer r.Close()
		code := 200
		categories := []*category{}
		var err interface{}
		if e == nil {
			for r.Next() {
				c := &category{}
				r.Scan(&c.ID, &c.Title, &c.Description)
				categories = append(categories, c)
			}
			if len(categories) == 0 {
				code = 404
				_, hasID := c.Request.URL.Query()["id"]
				if hasID {
					err = "Category not found"
				} else {
					err = "No categories"
				}
			}
		} else {
			err = string(e.Error())
			code = 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  categories,
		})
	}
}
