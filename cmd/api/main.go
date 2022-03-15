package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello GoCommerce")

	var cfg config

	flag.IntVar(&cfg.port, "port", 8081, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	router := gin.Default()
	router.GET("/v1/api-health-checker", app.healthCheckerHandler)

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLog.Printf("starting %v server at port %v", cfg.env, cfg.port)
	errorLog.Fatal(server.ListenAndServe())

}
