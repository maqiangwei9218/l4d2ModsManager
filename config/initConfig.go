package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func InitConfig() {
	cfg := GetConfig()

	m := map[string]bool{}
	for _, i := range GetAllItems() {
		i = strings.Replace(i, "\\", "/", -1)
		if cfg[i] {
			m[i] = true
		} else {
			m[i] = false
		}
	}

	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}

	// 写入文件
	err = os.WriteFile("config.json", jsonData, 0644)
}
