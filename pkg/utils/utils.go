package utils

import "github.com/google/uuid"

type Url_Config struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	Auth    string `json:"auth"`
	Payload string `json:"payload"`
}
type TaskRequest struct {
	TaskID    string       `json:"task_id"`
	TaskName  string       `json:"taskname"`
	Frequency int          `json:"frequency"`
	Config    []Url_Config `json:"config"`
}

func GenerateUUID() string {
	return uuid.New().String()
}
