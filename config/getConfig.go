package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Dirs  []string `yaml:"dirs"`
	Items []string `yaml:"items"`
}

func GetConfig() map[string]bool {
	var result map[string]bool

	data, err := os.ReadFile("config.json")
	if err != nil {
		return result
	}

	// 解析JSON到map
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result
	}

	return result
}
