package main

import (
	"clauses/database"
	"clauses/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
