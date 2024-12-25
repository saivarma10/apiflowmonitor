package main

import (
	api "apimonitor/internal/api"
	config "apimonitor/internal/config"
	db "apimonitor/internal/db"
	"apimonitor/internal/processor"
	logger "apimonitor/pkg/logger"
	"apimonitor/pkg/utils"
	"net/http"
)

func ping() string {
	return "pong"
}
func main() {
	_, err := config.ReadConfigFile("conf/config.json")
	if err != nil {
		panic(err)
	}
	err = db.Init()
	if err != nil {
		panic(err)
	}
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
	db.Close()
}
