package main

import (
	"log"

	"github.com/saladin2098/forum_api_gateway/api"
	"github.com/saladin2098/forum_api_gateway/api/handlers"
	"github.com/saladin2098/forum_api_gateway/config"
)

func main() {
	cfg := config.Load()
	h, err := handlers.NewHandler()
	if err != nil {
		log.Fatal(err)
	}
	
	r := api.NewGin(h)
	r.Run(cfg.HTTPPort)
}
