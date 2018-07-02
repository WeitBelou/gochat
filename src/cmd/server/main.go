package main

import (
	"log"

	"lib/api"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to read config: %+v", err)
	}
	log.Printf("config: %+v", cfg)

	r := gin.New()

	api.Register(r, api.Services{
		Auth: nil, // FIXME(i.kosolapov): Replace with actual implementation.
	})

	r.Run(cfg.Server.addr())
}
