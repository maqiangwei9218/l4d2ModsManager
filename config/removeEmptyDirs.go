package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveEmptyDirs(path string) error {
	// 先处理子目录
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fullPath := filepath.Join(path, entry.Name())
			err := RemoveEmptyDirs(fullPath)
			if err != nil {
				return err
			}
		}
	}

	// 检查当前目录是否为空
	entries, err = os.ReadDir(path)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Printf("Removing empty directory: %s\n", path)
		return os.Remove(path)
	}

	return nil
}
