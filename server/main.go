package main

import (
	"buycryptos/server/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	router.Start().Run()
}
