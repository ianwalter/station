package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/unrolled/logger"
)

type Handler struct {
	fs    http.Handler
	index []byte
	log   *logger.Logger
}

func containsDotFile(name string) bool {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "station/1.0.1")

	if containsDotFile(r.URL.Path) {
		w.WriteHeader(403)
	} else if path.Ext(r.URL.Path) != "" {
		// If the path has a file extension, use the FileServer to serve the static
		// file.
		h.fs.ServeHTTP(w, r)

	} else {
		// If the path doesn't have a file extension, serve the index.html file.
		w.Write(h.index)
	}
}

func main() {
	log := logger.New()

	// Determine the directory of static files to serve.
	dir, err := filepath.Abs("./")
	if len(os.Args) > 1 {
		dir, _ = filepath.Abs(os.Args[1])
	}

	// Read the index.html file and create the Handler instance.
	indexPath := path.Join(dir, "index.html")
	index, err := ioutil.ReadFile(indexPath)
	if err != nil {
		log.Printf("index.html not found: %s", indexPath)
	}
	handler := Handler{http.FileServer(http.Dir(dir)), index, log}

	// Determine what port to listen on.
	port := os.Getenv("PORT")
	if port == "" {
		port = "9876"
	}

	// Start the server.
	log.Println(fmt.Sprintf("🌐 Station up and running on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), log.Handler(handler)))
}
