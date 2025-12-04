package main

import (
	"log"

	"github.com/mesh-dell/blogging-platform-API/config"
	"github.com/mesh-dell/blogging-platform-API/internal/api"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	api.InitServer(config)
}
