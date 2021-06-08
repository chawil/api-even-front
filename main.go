package main

import (
	"embed"
	_ "embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed static
var staticEmbed embed.FS

func main() {
	fsys, err := fs.Sub(staticEmbed, "static")
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.StaticFS("/home", http.FS(fsys))

	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "/home") })

	r.GET("/api/even/:number", func(c *gin.Context) {
		result, err := strconv.Atoi(string(c.Param("number")))
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, struct {
			Result int       `json:"result"`
			IsEven bool      `json:"isEven"`
			Date   time.Time `json:"date"`
		}{Result: result, Date: time.Now().UTC()})
	})

	log.Println("Serving on http://localhost:8080/")
	err = r.Run()
	log.Fatalln(err)
}
