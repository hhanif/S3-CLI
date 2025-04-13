package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// config represents the configuration for the application
type Config struct {
	Minio       Minio  `yaml:"minio"`
	BucketName  string `yaml:"bucket_name"`
	FileContent string `yaml:"file_content"`
}

// Minio represents the configuration for the Minio server
type Minio struct {
	Server   string `yaml:"server"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSL      bool   `yaml:"ssl"`
}

// LoadConfig loads the configuration from the given file path
func LoadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	//validate config
	if cfg.BucketName == "" {
		return nil, fmt.Errorf("bucket name is required in config")
	}
	if cfg.Minio.User == "" {
		return nil, fmt.Errorf("minio user is required in config")
	}
	if cfg.Minio.Password == "" {
		return nil, fmt.Errorf("minio password is required in config")
	}
	return &cfg, nil
}
