package main

import (
	"yoyo-mall/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(g)

}
