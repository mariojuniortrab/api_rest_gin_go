package main

import (
	"github.com/mariojuniortrab/api-rest-gin-go/database"
	"github.com/mariojuniortrab/api-rest-gin-go/routes"
)

func main() {
	database.DatabaseConnect()
	routes.HandleRequests()
}
