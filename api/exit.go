package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleExit(c *gin.Context) {
	fmt.Println("exit")
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
	os.Exit(0)
}
