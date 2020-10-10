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

	base := "/api/v1"

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		ExposeHeaders: []string{"Content-Type", "Date"},
		// AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		AllowAllOrigins: true,
	}))

	// Items
	r.GET(base+"/items", item.GET())

	// Users
	r.GET(base+"/user", user.GET())

	return r
}
