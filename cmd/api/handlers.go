package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (app *application) healthCheckerHandler(c *gin.Context) {
	fmt.Fprintf(c.Writer, "environment: %v\n", app.config.env)
	fmt.Fprintf(c.Writer, "api version %v\n", version)
}
