package main

import (
	"github.com/Nivelian/codete-webscraping/api"
	"github.com/Nivelian/codete-webscraping/helpers"
)

func main() {
	helpers.InitLog()

	config := helpers.GetConfig()

	api.StartServer(config)
}
