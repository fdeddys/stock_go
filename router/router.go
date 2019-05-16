package router

import (
	"net/http"
	"stock/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g1 := r.Group("/api/master/groupbarang")
	{
		g1.GET("/id/:id", controller.FindGroupBarangByID)
		g1.GET("/code/:code", controller.FindGroupBarangByCode)
		g1.POST("", controller.AddGroupBarangByCode)
	}

	g2 := r.Group("/api/master/barang")
	{
		g2.GET("/id/:id", controller.FindBarangByID)
	}

	return r
}
