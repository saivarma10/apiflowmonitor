package utils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Url_Config struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	Auth    string `json:"auth"`
	Payload string `json:"payload"`
}
// type TaskRequest struct {
// 	TaskID    string       `json:"task_id"`
// 	TaskName  string       `json:"taskname"`
// 	Frequency int          `json:"frequency"`
// 	Config    []Url_Config `json:"config"`
// }

type UpdateRequest struct {
	TaskID    string       `json:"task_id"`
	Frequency int          `json:"frequency"`
}

type Task struct {
	TaskID string `json:"task_id"`
}

type APIConfig struct {
	URL               string `json:"url"`
	Method            string `json:"method"`
	Request  		  string `json:"request"`
}

type Dependency struct {
	APIIndex int    `json:"index"`
	APIKey   string	`json:"api_key"`
}

type TransactionAPI struct {
	URL         string      `json:"url"`
	Method      string      `json:"method"`
	Request     string      `json:"request"`
	Dependency map[string]Dependency `json:"dependency"`
}

type Transactions struct {
	TransactionID   string          `json:"transaction_id"`
	Name string          			`json:"name"`
	Frequency int          			`json:"frequency"`
	APIs []TransactionAPI 			`json:"apis"`
}

type TaskRequest struct {
	Transactions []Transactions `json:"transactions"`
}

func GenerateUUID() string {
	return uuid.New().String()
}
func GetKeyOrValueStartingWithDollar(m map[string]interface{}) (string, interface{}) {
	for key, value := range m {
		if strings.HasPrefix(key, "$") || strings.HasPrefix(fmt.Sprint(value), "$") {
			return key, value
		}
	}
	return "", nil
}
func SearchDynamicVariable(m map[string]interface{}, match string) (string, interface{}) {

		for key, value := range m {
			if strings.HasPrefix(key, match) || strings.HasPrefix(fmt.Sprint(value), match) {
				return key, value
			}
		}
	
	return "", nil
}
