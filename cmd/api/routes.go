package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) Routes() *gin.Engine {
	router := gin.Default()
	router.GET("/v1/api-health-checker", app.healthCheckerHandler)
	return router
}
