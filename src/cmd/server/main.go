package main

import (
	"log"
	"time"

	"lib/api"
	"lib/tokens"
	"lib/users"

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

	authService, err := users.New(cfg.Users)
	if err != nil {
		log.Panicf("failed to connect to create auth service")
	}
	api.Register(r, api.Services{
		Auth:   authService,
		Tokens: tokens.New(cfg.Tokens),
	})

	r.Run(cfg.Server.addr())
}
