package router

import (
	"buycryptos/server/src/item"
	"buycryptos/server/src/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Postgres driver
)

// Start the router
func Start() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	p := "/api/v1"
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// ExposeHeaders: []string{"Content-Type", "Date"},
		// AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		AllowAllOrigins: true,
	}))
	r.GET(p+"/items", item.GET())
	r.POST(p+"/item", item.POST())
	r.PUT(p+"/item", item.PUT())
	r.DELETE(p+"/item", item.DELETE())
	r.GET(p+"/reflinks", item.ReflinksGET())
	r.GET(p+"/featured", item.FeaturedGET())
	r.GET(p+"/categories", item.CategoriesGET())
	r.POST(p+"/categories", item.CategoryPOST())
	r.PUT(p+"/categories", item.CategoryPUT())
	r.DELETE(p+"/categories", item.CategoryDELETE())
	r.GET(p+"/ratings", item.RatingsGET())
	r.GET(p+"/rating", item.AverageRatingGET())
	r.POST(p+"/rating", item.RatingPOST())
	r.PUT(p+"/rating", item.RatingPUT())
	r.DELETE(p+"/rating", item.RatingDELETE())
	r.GET(p+"/featurings", item.FeaturingsGET())
	r.GET(p+"/auctions", item.AuctionsGET())
	r.GET(p+"/bids", item.BidsGET())
	r.GET(p+"/requests", item.RequestsGET())
	r.POST(p+"/requests", item.RequestPOST())
	r.GET(p+"/fundings", user.FundingsGET())
	r.GET(p+"/notifications", user.NotificationsGET())
	r.GET(p+"/notification", user.NotificationPOST())
	r.GET(p+"/notification", user.NotificationPUT())
	r.GET(p+"/users", user.GET())
	r.PUT(p+"/user", user.PUT())
	r.GET(p+"/profile", user.ProfileGET())
	r.PUT(p+"/password-update", user.PasswordUpdate())
	r.POST(p+"/user", user.POST())
	r.GET(p+"/login", user.Login())
	r.DELETE(p+"/user", user.DELETE())
	return r
}
