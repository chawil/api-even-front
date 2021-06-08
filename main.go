package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

//go:embed static
var static embed.FS

func main() {
	fsys, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatalln(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsys)))

	http.HandleFunc("/api/even", func(response http.ResponseWriter, req *http.Request) {
		defer (func(start time.Time) { log.Println("/api/even took", time.Since(start), "to respond") })(time.Now())
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}
		result, err := strconv.Atoi(string(body))
		if err != nil {
			log.Println(err)
		}

		err = json.NewEncoder(response).Encode(struct {
			Result int       `json:"result"`
			IsEven bool      `json:"isEven"`
			Date   time.Time `json:"date"`
		}{Result: result, Date: time.Now().UTC()})

		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Serving on http://localhost:8888/")

	err = http.ListenAndServe(":8888", nil)
	log.Fatalln(err)
}
