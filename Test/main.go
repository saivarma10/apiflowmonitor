package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestPayload struct {
	User     string `json:"user"`
	password string `json:"password"`
}
type ResponsePayload struct {
	User     string `json:"user"`
	password string `json:"password"`
}

type RequestPayload2 struct {
	User   string `json:"user"`
	exists bool   `json:"exists"`
}
type ResponsePayload2 struct {
	Sucess bool `json:"sucess"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestPayload
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}
	resp := ResponsePayload{
		User:     req.User,
		password: req.password,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
func apiHandler2(w http.ResponseWriter, r *http.Request) {
	var req RequestPayload2
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding request", http.StatusBadRequest)
		return
	}
	if req.User == "sai" {
		resp := ResponsePayload2{
			Sucess: true,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := ResponsePayload2{
		Sucess: false,
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(resp)
	fmt.Println("Error decoding request")
	return
}
func apiHandler3(w http.ResponseWriter, r *http.Request) {
	resp := ResponsePayload2{
		Sucess: true,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	// w.Write([]byte("Hello, World 3!"))
}

func main() {
	http.HandleFunc("/1", apiHandler)
	http.HandleFunc("/2", apiHandler2)
	http.HandleFunc("/3", apiHandler3)
	http.ListenAndServe(":9090", nil)

}

/*
the above program is a simple http server that listens on port 9090 and has 3 endpoints:
1. /1
2. /2
3. /3

i want to use this to test our other program that sends http requests to multiple endpoints

curl -X POST http://localhost:8081/create \                                                                                                                                ❮main|✚3...4
-H "Content-Type: application/json" \
-d '{
    "TaskID": "123456",
    "TaskName": "Updated Task Name2",
    "Frequency": 120,
    "Config": [
        {
            "Url": "http://localhost:9090/1",
            "Method": "POST",
            "Auth": "",
            "Payload": "{\"user\": \"sai\", \"password\": \"pass\"}"
        }, {
            "Url": "http://localhost:9090/2",
            "Method": "POST",
            "Auth": "",
            "Payload": "{\"user\": \"$user\", \"exists\": \"true\"}"
        }, {
            "Url": "http://localhost:9090/3",
            "Method": "GET"
        }
    ]
}'

*/
