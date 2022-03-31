package web

import (
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/peng225/wolper/service"
)

func Start(port int, addrAndPort, modulePath string) {
	router := gin.Default()

	router.LoadHTMLFiles(path.Join(modulePath, "html/index.html"))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/query", func(ctx *gin.Context) {
		key := ctx.Query("key")
		include := ctx.Query("include")
		exclude := ctx.Query("exclude")
		uniqStr := ctx.Query("unique")
		uniq := false
		if uniqStr == "on" {
			uniq = true
		}
		words := service.ClientQuery(addrAndPort,
			key, include, exclude, uniq)
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		responseBody := ""
		for _, word := range words {
			responseBody += word + "\n"
		}
		ctx.String(http.StatusOK, responseBody)
	})

	err := router.Run("localhost:" + strconv.Itoa(port))
	if err != nil {
		panic("Failed to start server.")
	}
}
