package main

import (
	"entry-exit-api/database"
	"entry-exit-api/routes"
)

func main() {

	err := database.InitDB()
	if err != nil {
    panic("failed to initialize database")
  }

	routes.Init()
}
