package main

import (
	api "apimonitor/internal/api"
	logger "apimonitor/pkg/logger"
	db "apimonitor/internal/db"
	config "apimonitor/internal/config"
	"net/http"
)

func ping() string {
	return "pong"
}
func main() {
	_, err := config.ReadConfigFile("../conf/config.json")
	if err != nil {
		panic(err)
	}
	err = db.Init()
	if err !=nil{
		panic(err)
	}
	println(ping())
	log := logger.GetLogger()
	// resp := processor.CurlGet("https://www.google.com")
	// println(resp)
	router := api.SetupRoutes()
	log.Println("Server is starting on port 8081...")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Info().Msgf("Error starting server: %v", err)
	}
	db.Close()
}
