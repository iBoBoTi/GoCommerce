package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) healthCheckerHandler(c *gin.Context) {
	reqStruct := struct {
		Status  string `json:"status"`
		Env     string `json:"env"`
		Version string `json:"version"`
	}{
		Status:  "available",
		Env:     app.config.env,
		Version: version,
	}
	js, err := json.Marshal(reqStruct)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	js = append(js, '\n')
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(js)
}
