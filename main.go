package main

import (
	"log"
	"os"

	"github.com/djaustin/urn-lookup/urn"
	"github.com/gin-gonic/gin"
)

func main() {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	urn := urn.NewController(redisAddress)

	r.GET("/urn/:id", urn.GetURN)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
