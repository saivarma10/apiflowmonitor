package api

import (
	"apimonitor/internal/db"
	processor "apimonitor/internal/processor"
	scheduler "apimonitor/internal/scheduler"
	"apimonitor/pkg/utils"
	util "apimonitor/pkg/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
	"time"
)

var tm = scheduler.NewTaskManager()

func addTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to add task.")
	var req util.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	for _,transaction := range req.Transactions{
		fmt.Println("Executing transaction ",transaction.TransactionID)
		if transaction.Frequency <= 0 {
			http.Error(w, fmt.Sprintf("Frequency must be greater than 0"), http.StatusBadRequest)
			continue
		}

		err := tm.CreateTask(transaction.TransactionID, transaction.Name, time.Duration(transaction.Frequency)*time.Second, processor.CurlRun(&transaction,transaction.TransactionID))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating task: %v", err), http.StatusInternalServerError)
			return
		} else {
			fmt.Println("Task created successfully")
		}
	
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task added successfully"))
	}
	
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to update task.")
	var req util.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	
	if req.Frequency <= 0 {
		http.Error(w, fmt.Sprintf("Frequency must be greater than 0"), http.StatusBadRequest)
		return
	}
	err := tm.UpdateTask(req.TaskID, time.Duration(req.Frequency)*time.Second)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating task: %v", err), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Task updated successfully")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task updated successfully"))
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to delete task.")
	var req util.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	err := tm.DeleteTask(req.TaskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting task: %v", err), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Task deleted successfully")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
}
func GetTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to get task.")
	var req util.Task
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	task, err := tm.GetTask(req.TaskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting task: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Task: %+v", task)))
}


func GetTransaction(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received request to get transaction.")
	var req utils.GetTransaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	transaction, err := strconv.Atoi(req.TransactionID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting transaction ID to integer: %v", err), http.StatusBadRequest)
		return
	}
	response, err := db.GetTransaction(transaction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting transaction: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Transaction: %+v", response)))
}