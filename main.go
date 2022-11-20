package main

import (
	"assignment-api/database"
	"assignment-api/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
