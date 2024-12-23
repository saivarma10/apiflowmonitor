package api

import (
	processor "apimonitor/internal/processor"
	scheduler "apimonitor/internal/scheduler"
	util "apimonitor/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var tm = scheduler.NewTaskManager()

func addTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to add task. %v")
	var req util.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("request is ", req)
	if req.Frequency <= 0 {
		http.Error(w, fmt.Sprintf("Frequency must be greater than 0"), http.StatusBadRequest)
		return
	}
	err := tm.CreateTask(req.TaskID, req.TaskName, time.Duration(req.Frequency)*time.Second, processor.CurlRun(req.Config))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating task: %v", err), http.StatusInternalServerError)
		return
	} else {
		fmt.Println("Task created successfully")
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task added successfully"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received request to update task.")
	var req util.TaskRequest
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
	var req util.TaskRequest
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
	var req util.TaskRequest
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
