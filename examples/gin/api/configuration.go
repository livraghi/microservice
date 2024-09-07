package main

import (
	"fmt"
	"github.com/livraghi/microservice/configuration"
)

//github.com/livraghi/microservice/configuration

func LoadConfiguration() (*Configuration, error) {
	var cfg Configuration
	err := configuration.ReadConfiguration(&cfg)
	if err != nil {
		fmt.Println("error loading configuration")
		return nil, err
	}
	fmt.Printf("configuration loaded: %v\n", cfg)
	return &cfg, nil
}

type Configuration struct {
}
