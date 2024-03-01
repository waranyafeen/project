package main

import (
	"github.com/waranyafeen/project/routers"
	"github.com/waranyafeen/project/entity"
)

func main() {
	entity.SetupDatabase("WaranratDB")
	entity.SetupData(entity.DB())
	route := routers.SetupRouter()
	routers.InitRouter(route)
	route.Run()

}