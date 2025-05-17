package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {

	Database struct{
		Host string `yaml:"host"`
		Port int `yaml:"port"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		DBName string `yaml:"dbname"`
		SSLMmode string `yaml:"sslmode"`
	} `yaml:"database"`
}

func LoadConfig(path string) (*Config, error){
	config := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config file error: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("yaml decode error: %w", err)
	}
	return config, nil
}