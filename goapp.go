package main

import (
	"golang-mock/database"
	"golang-mock/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
