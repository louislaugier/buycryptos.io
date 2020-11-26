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
	r.GET(p+"/categories", item.CategoriesGET())
	r.GET(p+"/ratings", item.RatingsGET())
	r.GET(p+"/featurings", item.FeaturingsGET())
	r.GET(p+"/auctions", item.AuctionsGET())
	r.GET(p+"/bids", item.BidsGET())
	r.GET(p+"/requests", item.RequestsGET())
	r.POST(p+"/requests", item.RequestPOST())
	r.GET(p+"/fundings", user.FundingsGET())
	r.GET(p+"/notifications", user.NotificationsGET())
	r.GET(p+"/users", user.GET())
	r.POST(p+"/users", user.POST())
	return r
}
