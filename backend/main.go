package main

import (
	"github.com/waranyafeen/project/backend/routers"
	"github.com/waranyafeen/project/backend/entity"
)

func main() {
	entity.SetupDatabase("PJWaranratDB")
	entity.SetupData(entity.DB())
	route := routers.SetupRouter()

	// init Routes
	routers.InitRouter(route)

	// Run the server
	route.Run()
}

