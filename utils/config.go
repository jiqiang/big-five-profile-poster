package utils

import (
	"encoding/json"
	"io/ioutil"
)

// Config is a struct hosting config data.
type Config struct {
	Endpoint string
	Email    string
	Source   string
}

// GetConfig returns a Config struct which hosts config data.
func GetConfig(path string) Config {
	configBytes := GetFileContent(path)
	config := Config{}
	err := json.Unmarshal(configBytes, &config)
	if err != nil {
		panic(err)
	}

	return config
}

// GetFileContent returns file content in byte array format.
func GetFileContent(path string) []byte {
	byt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return byt
}
