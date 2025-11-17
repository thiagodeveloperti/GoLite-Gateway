package config

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Path     string   `yaml:"path" json:"path"`
	Upstream string   `yaml:"upstream" json:"upstream"`
	Methods  []string `yaml:"methods" json:"methods"`
	Auth     bool     `yaml:"auth" json:"auth"`
}

type JWT struct {
	Secret string `yaml:"secret" json:"secret"`
	Issuer string `yaml:"issuer" json:"issuer"`
}

type Config struct {
	Routes []Route `yaml:"routes"`
	JWT    JWT     `yaml:"jwt"`
}

var Current Config

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	Current = cfg
	j, _ := json.MarshalIndent(cfg, "", "  ")
	println("Loaded config:\n", string(j))

	return cfg, nil
}
