package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type Config struct {
	Region          string `json:"region"`
	ContainerName   string `json:"container_name"`
	ObjectName      string `json:"object_name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	SaveLocation    string `json:"save_location"`
	ExtractLocation string `json:"extract_location"`
}
func LoadConfig(file string) *Config {
	fmt.Println("Reading file: ", file)
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}


	var config *Config
	err = json.Unmarshal([]byte(raw), &config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return config
}
