package main

import (
	"buycryptos.io/server/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	router.Start().Run()
}
