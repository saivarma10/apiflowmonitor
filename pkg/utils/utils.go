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
type TaskRequest struct {
	TaskID    string       `json:"task_id"`
	TaskName  string       `json:"taskname"`
	Frequency int          `json:"frequency"`
	Config    []Url_Config `json:"config"`
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
func SearchDynamicVariable(m []map[string]interface{}, match string) (string, interface{}) {
	for _, v := range m {
		for key, value := range v {
			if strings.HasPrefix(key, match) || strings.HasPrefix(fmt.Sprint(value), match) {
				return key, value
			}
		}
	}
	return "", nil
}
