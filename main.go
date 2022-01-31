package main

import (
	"yoyo-mall/model"
	"yoyo-mall/pkg/config"
	"yoyo-mall/pkg/log"
	"yoyo-mall/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// logger sync
	defer log.SyncLogger()

	// init DB
	model.DB.Init()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(g)

	log.Info(g.Run(viper.GetString("addr")).Error())
}
