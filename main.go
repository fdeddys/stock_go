package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "stock/db"

	"github.com/gin-gonic/gin"

	r "stock/router"
)

var db = make(map[string]string)

func main() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)

	router := r.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
