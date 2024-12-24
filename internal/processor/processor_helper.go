package processor

import (
	logger "apimonitor/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

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
		for _, c := range config {
			args := []string{
				c.Url,
				c.Method,
				c.Auth,
				c.Payload,
			}
			resp, err := runTransactionMonitorBinary(args)
			fmt.Println("Response and error: ", resp, err)
			//store reponse to mysql
		}
	}
}
