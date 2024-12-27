package main

import (
	api "apimonitor/internal/api"
	db "apimonitor/internal/db"
	logger "apimonitor/pkg/logger"
	"net/http"
)

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	log := logger.GetLogger()
	router := api.SetupRoutes()
	log.Println("Server is starting on port 8081...")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Info().Msgf("Error starting server: %v", err)
	}
	db.Close()

	// config := []utils.Url_Config{
	// 	{
	// 		Url:     "https://reqres.in/api/users",
	// 		Method:  "POST",
	// 		Auth:    "",
	// 		Payload: "{\"name\": \"morpheus\", \"job\": \"leader\"}",
	// 	},
	// }
	// task := processor.CurlRun(config,"")
	// task()
}
