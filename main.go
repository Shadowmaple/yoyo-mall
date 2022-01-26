package main

import (
	"log"
	"yoyo-mall/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set gin mode.
	gin.SetMode("debug")

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(g)

	log.Println(g.Run(":4096"))
}
