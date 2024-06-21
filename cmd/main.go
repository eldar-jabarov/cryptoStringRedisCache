package main

import (
	"cryptoStringRedisCache/lib/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	app.New().
		Run().
		Wait()
}
