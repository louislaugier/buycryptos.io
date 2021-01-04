package item

import (
	"github.com/louislaugier/buycryptos/server/database"
	"encoding/json"
	"os"

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
		code, categories := 200, []*category{}
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
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": err,
			"data":  categories,
		})
	}
}

// CategoryPOST export
func CategoryPOST() func(c *gin.Context) {
	return func(c *gin.Context) {
		t, cg, code := c.Request.URL.Query()["token"][0], &category{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &cg)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx.Exec("insert into categories (title, description) values ($1,$2);", &cg.Title, &cg.Description)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &cg,
		})
	}
}

// CategoryPUT export
func CategoryPUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		t, ID, cg, code := q["token"][0], q["id"][0], &category{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &cg)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx.Exec("update categories set title=$1, description=$2 where id='"+ID+"';", &cg.Title, &cg.Description)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &cg,
		})
	}
}

// CategoryDELETE export
func CategoryDELETE() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		t, ID, cg, code := q["token"][0], q["id"][0], &category{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &cg)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx.Exec("delete from categories where id='" + ID + "';")
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &cg,
		})
	}
}
