package api

import (
	"fmt"
	"go-example/config"
	"go-example/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Body_ struct {
	DirPath string `json:"dirPath" binding:"required"`
	Type    string `json:"type" binding:"required,oneof=add remove"`
}

func HandleMod(c *gin.Context) {
	var body Body_
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entries, err := os.ReadDir(body.DirPath)
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
	}

	errors_remove := []string{}

	m := map[string]string{}

	for _, entry := range entries {
		if !entry.IsDir() {
			if strings.HasSuffix(strings.ToLower(entry.Name()), ".vpk") {
				from := filepath.Join(body.DirPath, entry.Name())
				to := utils.FormatFilenameFromPath(from)
				m[from] = to
			}
		}
	}

	if body.Type == "add" {
		config.SetConfig(body.DirPath, true)
		for from, to := range m {
			utils.CopyFile(from, to)
		}
	}

	if body.Type == "remove" {
		for _, to := range m {
			err = os.Remove(to)
			if err != nil {
				errors_remove = append(errors_remove, err.Error())
			}
		}
		if len(errors_remove) == 0 {
			config.SetConfig(body.DirPath, false)
		}
	}

	if len(errors_remove) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"errors": errors_remove,
		})
	}

}
