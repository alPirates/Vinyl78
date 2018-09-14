package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	filepath = "config/config.json"
)

type Config struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func createDefautlConfig() error {
	data, err := json.MarshalIndent(&Config{
		Email:    "test@test.com",
		Password: "test_password",
	}, "", "\t")

	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (conf *Config) ReadConfig() error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("Config file not found, creating....")
		createDefautlConfig()
		return err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	return nil
}
