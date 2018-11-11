package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	index, _ := ioutil.ReadFile("README.md")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if path.Ext(r.URL.Path) != "" {
			fs.ServeHTTP(w, r)
		} else {
			w.Write(index)
		}
	})

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
