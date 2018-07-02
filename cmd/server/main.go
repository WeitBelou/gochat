package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to read config: %+v", err)
	}
	log.Printf("config: %+v", cfg)

	r := gin.New()
	r.Run(":8080")
}
