package main

import (
	api "apimonitor/internal/api"
	processor "apimonitor/internal/processor"
	logger "apimonitor/pkg/logger"
	"net/http"
)

func ping() string {
	return "pong"
}
func main() {
	println(ping())
	log := logger.GetLogger()
	resp := processor.CurlGet("https://www.google.com")
	println(resp)
	router := api.SetupRoutes()
	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Info().Msgf("Error starting server: %v", err)
	}
}
