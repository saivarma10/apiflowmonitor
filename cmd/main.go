package main

import (
	api "apimonitor/internal/api"
	logger "apimonitor/pkg/logger"
	"net/http"
)

func ping() string {
	return "pong"
}
func main() {
	println(ping())
	log := logger.GetLogger()
	// resp := processor.CurlGet("https://www.google.com")
	// println(resp)
	router := api.SetupRoutes()
	log.Println("Server is starting on port 8081...")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Info().Msgf("Error starting server: %v", err)
	}
}
