package main

import (
	api "apimonitor/internal/api"
	"apimonitor/internal/processor"
	logger "apimonitor/pkg/logger"
	"apimonitor/pkg/utils"
	"net/http"
)

func ping() string {
	return "pong"
}
func main() {
	println(ping())
	config := []utils.Url_Config{
		{
			Url:     "https://reqres.in/api/users",
			Method:  "POST",
			Auth:    "",
			Payload: "{\"name\": \"morpheus\", \"job\": \"leader\"}",
		},
	}
	task := processor.CurlRun(config)
	task()

	log := logger.GetLogger()
	router := api.SetupRoutes()
	log.Println("Server is starting on port 8081...")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Info().Msgf("Error starting server: %v", err)
	}

}
