package main

import (
	"log"
	"time"

	"lib/api"
	"lib/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	// FIXME(i.kosolapov): Wait for database
	time.Sleep(5 * time.Second)

	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to read config: %+v", err)
	}
	log.Printf("config: %+v", cfg)

	r := gin.New()

	authService, err := auth.New(cfg.Auth)
	if err != nil {
		log.Panicf("failed to connect to create auth service")
	}

	api.Register(r, api.Services{
		Auth: authService,
	})

	r.Run(cfg.Server.addr())
}
