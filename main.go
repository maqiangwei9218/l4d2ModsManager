package main

import (
	"embed"
	"fmt"
	"go-example/api"
	"go-example/config"
	"go-example/utils"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed ui/dist/*
var staticFiles embed.FS

func main() {
	config.InitConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(cors.Default())

	r.GET("/mods", api.HandleMods)
	r.POST("/mod", api.HandleMod)

	subFS, err := fs.Sub(staticFiles, "ui/dist")
	if err != nil {
		return
	}
	r.StaticFS("/front", http.FS(subFS))
	utils.OpenBrowser("http://127.0.0.1:9090/front/")
	fmt.Println("http://127.0.0.1:9090/front/")
	r.Run(":9090")
}
