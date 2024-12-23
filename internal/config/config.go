package config


import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)

var conf *ConfigFile

func GetConfig() *ConfigFile {
	return conf
}

func ReadConfigFile(filename string) (*ConfigFile, error) {
	fmt.Printf("Reading config file: %s\n", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	conf = new(ConfigFile)
	if err := json.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	fmt.Printf("Config Loaded : %+v\n", conf)

	return conf, nil
}


