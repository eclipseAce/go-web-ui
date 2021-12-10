package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed ui/dist
	embedFiles embed.FS
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ui/")
	})

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"data": "pong"})
		})
	}

	uiDistFiles, err := fs.Sub(embedFiles, "ui/dist")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/ui", http.FS(uiDistFiles))

	r.Run(":3000")
}
