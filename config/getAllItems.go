package config

import (
	"os"
	"path/filepath"
)

func hasFilesButNoSubdirs1(dirPath string) bool {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false
	}

	hasFiles := false
	hasSubdirs := false

	for _, entry := range entries {
		if entry.IsDir() {
			hasSubdirs = true
			break
		} else {
			hasFiles = true
		}
	}

	return hasFiles && !hasSubdirs
}

func GetAllItems() []string {
	var dirs []string
	err := filepath.WalkDir("custommods", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if hasFilesButNoSubdirs1(path) {
				dirs = append(dirs, path)
			}
		}
		return nil
	})

	if err != nil {
		return dirs
	}
	return dirs
}
