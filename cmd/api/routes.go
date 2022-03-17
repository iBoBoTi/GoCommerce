package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) Routes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/api-health-checker", app.healthCheckerHandler)
	return router
}
