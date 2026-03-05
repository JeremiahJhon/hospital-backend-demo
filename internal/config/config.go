package config

import "os"

type Config struct {
	HISBaseURL string
}

func LoadConfig() *Config {
	return &Config{
		HISBaseURL: os.Getenv("HIS_BASE_URL"),
	}
}
