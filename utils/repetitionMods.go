package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// 检测目录下重复的文件
// 根据文件名判断
func RepetitionMods(dir string, suffixes []string) []string {
	var m = map[string]int{}
	var r = []string{}

	filepath.WalkDir("custommods", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {

			for _, v := range suffixes {
				if strings.HasSuffix(strings.ToLower(path), v) {
					var filename = filepath.Base(path)
					var count = m[filename]
					m[filename] = count + 1
				}
			}
		}

		return nil
	})

	for k, v := range m {
		if v >= 2 {
			r = append(r, k)
		}
	}

	return r
}
