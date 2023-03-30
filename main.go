package main

import (
	"book-challenge/database"
	"book-challenge/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
