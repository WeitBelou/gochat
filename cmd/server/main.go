package main

import (
	"log"

	"gochat/lib/api"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to read config: %+v", err)
	}
	log.Printf("config: %+v", cfg)

	r := gin.New()
	api.Register(r)

	r.Run(cfg.Server.addr())
}
