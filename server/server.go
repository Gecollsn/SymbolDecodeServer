package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SymparseRun() {
	fmt.Println("hello server")
	r := gin.Default()
	r.LoadHTMLGlob("web/dist/*.html")
	r.Static("/assets", "./web/dist/assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.Run(":7000") // 监听并在 0.0.0.0:8080 上启动服务
}
