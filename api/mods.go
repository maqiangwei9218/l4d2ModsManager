package api

import (
	"go-example/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mod struct {
	Status         int    `json:"status"`
	DirPath        string `json:"dirPath"`
	TargetFileName string `json:"targetFileName"`
}

func HandleMods(c *gin.Context) {
	// data := []Mod{}
	// items := config.GetAllItems()
	// fmt.Println("items", items)
	// for _, item := range items {
	// 	data = append(data, Mod{Status: 1, DirPath: item,TargetFileName:utils.FormatFilenameFromPath(item)})
	// }
	c.JSON(http.StatusOK, config.GetConfig())
}
