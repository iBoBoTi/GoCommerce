package main

import (
	"fmt"
	"github.com/iBoBoTi/go-commerce/internal/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// load .env
	err := godotenv.Load("go-commerce.env")
	if err != nil {
		log.Fatal("Error loading go-commerce.env file")
		return
	}
	fmt.Println("Hello GoCommerce")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// set application config
	var cfg config
	cfg.port = os.Getenv("SERVER_PORT")
	if cfg.port == "" {
		cfg.port = "8081"
	}
	cfg.env = os.Getenv("ENV")
	cfg.db = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.port),
		Handler:      app.Routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Connect to the database and set the database instance on the application
	infoLog.Println("Connecting to database...")
	goCommerceRepo, err := repository.NewGoCommerceRepo(cfg.db)
	if err != nil {
		errorLog.Fatal(err)
	}
	app.db = goCommerceRepo.DB

	infoLog.Printf("starting %v server at port %v", cfg.env, cfg.port)
	errorLog.Fatal(server.ListenAndServe())

}
