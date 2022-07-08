package main

import (
	"errors"
	"github.com/chew01/verse-api/db"
	"github.com/chew01/verse-api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		err = errors.New("error loading .env: " + err.Error())
		panic(err)
	}

	err = db.Init()
	if err != nil {
		err = errors.New("error initializing database pool: " + err.Error())
		panic(err)
	}

	router := routes.New()
	err = router.Run("localhost:8080")
	if err != nil {
		err = errors.New("error starting server: " + err.Error())
		panic(err)
	}
}
