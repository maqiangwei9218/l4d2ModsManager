package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func SetConfig(key string, value bool) {
	cfg := GetConfig()
	cfg[key] = value
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}

	err = os.WriteFile("config.json", jsonData, 0644)
}
