package main

import (
	"github.com/kristiansantos/go-api-rest/api-web/routes"
	"github.com/kristiansantos/go-api-rest/config"
)

func main() {
	config.Start()
	routes.HandleRequests()
}
