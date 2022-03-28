package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello GoCommerce")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	var cfg config
	cfg.port = os.Getenv("SERVER_PORT")
	if cfg.port == "" {
		cfg.port = "8080"
	}
	cfg.env = os.Getenv("ENV")
	//cfg.db = fmt.Sprintf("%v://%v@%v:%v/%v?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
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

	infoLog.Printf("starting %v server at port %v", cfg.env, cfg.port)
	errorLog.Fatal(server.ListenAndServe())

}

func PostgresConnection(cfg config) (*sql.DB, error) {
	//db, err := pgx.Connect(context.Background(), cfg.db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	db, err := sql.Open("postgres", cfg.db)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(15 * time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
