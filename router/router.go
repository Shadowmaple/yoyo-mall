package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const BasePath = "api/v1"

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	// g.Use(middleware.NoCache)
	// g.Use(middleware.Options)
	// g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	return g
}
