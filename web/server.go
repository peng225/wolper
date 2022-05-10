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
		entropySortStr := ctx.Query("entropy sort")
		uniq := false
		if uniqStr == "on" {
			uniq = true
		}
		entropySort := false
		if entropySortStr == "on" {
			entropySort = true
		}
		words := service.ClientQuery(addrAndPort,
			key, include, exclude, uniq, entropySort)
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		responseBody := ""
		if len(words) == 0 {
			responseBody = "No results found."
		} else {
			for _, word := range words {
				responseBody += word + "\n"
			}
		}
		ctx.String(http.StatusOK, responseBody)
	})

	err := router.Run("0.0.0.0:" + strconv.Itoa(port))
	if err != nil {
		panic("Failed to start server.")
	}
}
