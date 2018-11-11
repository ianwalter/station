package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type Handler struct {
	fs    http.Handler
	index []byte
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if path.Ext(r.URL.Path) != "" {
		// If the path has a file extension, use the FileServer to serve the static
		// file.
		h.fs.ServeHTTP(w, r)
	} else {
		// If the path doesn't have a file extension, serve the index.html file.
		w.Write(h.index)
	}
}

func main() {
	// Determine the directory of static files to serve.
	dir, err := filepath.Abs("./")
	if len(os.Args) > 1 {
		if dir, err = filepath.Abs(os.Args[1]); err != nil {
			log.Println(fmt.Sprintf("Path not found: %s", os.Args[1]))
			os.Exit(1)
		}
	}

	// Read the index.html file and create the Handler instance.
	indexPath := path.Join(dir, "index.html")
	index, err := ioutil.ReadFile(indexPath)
	if err != nil {
		log.Println(fmt.Sprintf("index.html not found: %s", indexPath))
		os.Exit(1)
	}
	handler := Handler{http.FileServer(http.Dir(dir)), index}

	// Determine what port to listen on.
	port := os.Getenv("PORT")
	if port == "" {
		port = "9876"
	}

	// Start the server.
	log.Println("Listening...")
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
