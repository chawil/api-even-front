package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
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

	backendUrl := os.Getenv("BACKEND_URL")
	log.Printf("BACKEND_URL: %#+v\n", backendUrl)

	r := gin.Default()

	r.StaticFS("/home", http.FS(fsys))

	r.GET("/", func(c *gin.Context) { c.Redirect(http.StatusTemporaryRedirect, "/home") })

	r.GET("/api/even/:number", func(c *gin.Context) {
		numberString := c.Param("number")
		result, err := strconv.Atoi(string(numberString))
		if err != nil {
			log.Println(err)
		}

		resp, err := http.Get(backendUrl + "/even/" + numberString)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		backendResult := struct {
			IsEven bool `json:"isEven"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&backendResult)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, struct {
			Result    int       `json:"result"`
			IsEven    bool      `json:"isEven"`
			Date      time.Time `json:"date"`
			Namespace string    `json:"namespace"`
		}{Result: result, Date: time.Now().UTC(), IsEven: backendResult.IsEven, Namespace: os.Getenv("NAMESPACE")})
	})

	log.Println("Serving on http://localhost:8080/")
	err = r.Run()
	log.Fatalln(err)
}
