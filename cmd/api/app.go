package main

import "log"

const version = "1.0.0"

//config to hold all the configuration of the application
type config struct {
	port int
	env  string
}

//application to hold all the dependencies for the http handlers, helpers and middleware
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}
