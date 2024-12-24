package processor

import (
	logger "apimonitor/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	util "apimonitor/pkg/utils"
	// curl "github.com/andelf/go-curl"
)

var log = logger.GetLogger()

type Response struct {
	StatusCode   int                    `json:"status_code"`
	ResponseTime float64                `json:"response_time"`
	ResponseData map[string]interface{} `json:"response_data"`
}

func runTransactionMonitorBinary(args []string) (*Response, error) {
	binaryPath := "./cmd/c_program/assets/transaction_monitor"
	fmt.Printf("args is %v\n", args)

	cmd := exec.Command(binaryPath, args...)
	fmt.Println("cmd is ", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start command: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("failed to execute binary: %w", err)
	}

	if errOut.Len() > 0 {
		return nil, fmt.Errorf("stderr: %s", errOut.String())
	}
	var response Response
	err = json.Unmarshal(out.Bytes(), &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}
	fmt.Println(response)
	return &response, nil
}

func CurlRun(config []util.Url_Config) func() {
	return func() {
		var resp2 []*Response
		var TotalResponseData []map[string]interface{}
		var err error
		for _, c := range config {
			// var resp *Response
			var jsonMap map[string]interface{}
			json.Unmarshal([]byte(c.Payload), &jsonMap)
			fmt.Println("jsonMap is ", jsonMap)
			dynamicVar, _ := util.GetKeyOrValueStartingWithDollar(jsonMap)
			fmt.Println("dynamicVar is ", dynamicVar)
			if dynamicVar != "" {
				k, v := util.SearchDynamicVariable(TotalResponseData, dynamicVar)
				fmt.Println("k is ", k)
				fmt.Println("v is ", v)
				// we have to replace this key and value  in c.Payload
				if k == "" {
					k, _ = v.(string)
				}
				c.Payload = strings.Replace(c.Payload, dynamicVar, fmt.Sprint(k), -1)
			}
			args := []string{
				c.Url,
				c.Method,
				c.Auth,
				c.Payload,
			}
			resp, err := runTransactionMonitorBinary(args)
			fmt.Println("Response and error: ", resp, err)
			TotalResponseData = append(TotalResponseData, resp.ResponseData)
			resp2 = append(resp2, resp)
			//store reponse to mysql
		}
		for i := range resp2 {
			fmt.Printf("All Responses:%v err %v", *resp2[i], err)
		}

	}
}
