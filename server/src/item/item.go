package item

import (
	"encoding/json"
	"os"
	"time"

	"buycryptos/server/database"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID                    *int      `json:"id"`
	Title                 *string   `json:"title"`
	BaseLink              *string   `json:"base_link"`
	RefLink               *string   `json:"ref_link"`
	Description           *string   `json:"description"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	RefLinkOwnerUserEmail *int      `json:"ref_link_owner_user_email"`
	FeaturerUserEmail     *string   `json:"featurer_user_email"`
	CategoryID            *int      `json:"category_id"`
	ImagePath             *string   `json:"image_path"`
}

// GET export
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q, cID := c.Request.URL.Query(), ""
		if _, i := q["category_id"]; i {
			cID = "where category_id='" + q["category_id"][0] + "' "
		}
		r, e := database.DB.Query("select id, title, base_link, ref_link, description, created_at, updated_at, ref_link_owner_user_email, featurer_user_email, image_path from items " + cID + database.StandardizeQuery(q) + ";")
		defer r.Close()
		code, items := 200, []*item{}
		var err interface{}
		if e == nil {
			for r.Next() {
				i := &item{}
				r.Scan(&i.ID, &i.Title, &i.BaseLink, &i.RefLink, &i.Description, &i.CreatedAt, &i.UpdatedAt, &i.RefLinkOwnerUserEmail, &i.FeaturerUserEmail, &i.ImagePath)
				items = append(items, i)
			}
			if len(items) == 0 {
				code = 404
				_, hasID := q["id"]
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
			"error": &err,
			"data":  &items,
		})
	}
}

// ReflinksGET export
func ReflinksGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		rl, code, items := q["ref_link_owner_email"][0], 200, []*item{}
		r, e := database.DB.Query("select id, title, base_link, ref_link, description, created_at, updated_at, ref_link_owner_user_email, featurer_user_email, image_path from items where ref_link_owner_email='" + rl + "' " + database.StandardizeQuery(q) + ";")
		defer r.Close()
		var err interface{}
		if e == nil {
			for r.Next() {
				i := &item{}
				r.Scan(&i.ID, &i.Title, &i.BaseLink, &i.RefLink, &i.Description, &i.CreatedAt, &i.UpdatedAt, &i.RefLinkOwnerUserEmail, &i.FeaturerUserEmail, &i.ImagePath)
				items = append(items, i)
			}
			if len(items) == 0 {
				code = 404
				err = "No reflinks"
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &items,
		})
	}
}

// FeaturedGET export
func FeaturedGET() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		ID, em, code, items := q["id"][0], q["featurer_user_email"][0], 200, []*item{}
		r1, e := database.DB.Query("select email from users where id='" + ID + "';")
		defer r1.Close()
		for r1.Next() {
			r1.Scan(&em)
		}
		var err interface{}
		if e == nil {
			r2, _ := database.DB.Query("select id, title, base_link, ref_link, description, created_at, updated_at, ref_link_owner_user_email, featurer_user_email, image_path from items where featurer_user_email='" + em + "' " + database.StandardizeQuery(q) + ";")
			defer r2.Close()
			for r2.Next() {
				i := &item{}
				r2.Scan(&i.ID, &i.Title, &i.BaseLink, &i.RefLink, &i.Description, &i.CreatedAt, &i.UpdatedAt, &i.RefLinkOwnerUserEmail, &i.FeaturerUserEmail, &i.ImagePath)
				items = append(items, i)
			}
			if len(items) == 0 {
				code, err = 404, "No featured items"
			}
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &items,
		})
	}
}

// POST export
func POST() func(c *gin.Context) {
	return func(c *gin.Context) {
		t := c.Request.URL.Query()["token"][0]
		i, code := &item{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &i)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx.Exec("insert into items (title, base_link, ref_link, description, category_id, image_path) values ($1,$2,$3,$4,$5,$6);", &i.Title, &i.BaseLink, &i.RefLink, &i.Description, &i.CategoryID, &i.ImagePath)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &i,
		})
	}
}

// PUT export
func PUT() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		t, ID := q["token"][0], q["id"][0]
		i, code := &item{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &i)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx.Exec("update items set title=$1, base_link=$2, ref_link=$3, description=$4, updated_at=$5, ref_link_owner_email=$6, featurer_user_email=$7, category_id=$8, image_path=$9 where id='"+ID+"';", &i.Title, &i.BaseLink, &i.RefLink, &i.Description, &i.UpdatedAt, &i.UpdatedAt, &i.RefLinkOwnerUserEmail, &i.FeaturerUserEmail, &i.CategoryID, &i.ImagePath)
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &i,
		})
	}
}

// DELETE export
func DELETE() func(c *gin.Context) {
	return func(c *gin.Context) {
		q := c.Request.URL.Query()
		t, ID := q["token"][0], q["id"][0]
		i, code := &item{}, 200
		p, _ := c.GetRawData()
		json.Unmarshal(p, &i)
		tx, e := database.DB.Begin()
		var err interface{}
		if e == nil && t == os.Getenv("ADMIN_TOKEN") {
			_, e = tx.Exec("delete from items where id='" + ID + "';")
			tx.Commit()
		} else {
			err, code = string(e.Error()), 500
		}
		c.JSON(code, &gin.H{
			"error": &err,
			"data":  &i,
		})
	}
}
