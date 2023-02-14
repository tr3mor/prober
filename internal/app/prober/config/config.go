package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Period         time.Duration `yaml:"period"`
	Timeout        time.Duration `yaml:"timeout"`
	PrometheusPort int           `yaml:"prometheus_port"`
	Targets        []Target      `yaml:"targets"`
}

type Target struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func ParseConfig(path string, log *zap.SugaredLogger) *Config {
	var cfg Config
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Failed to parse config file: %s", err)
	}
	return &cfg
}
