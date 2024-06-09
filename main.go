package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/", logger(http.HandlerFunc(staticHandler)))
	router.Handle("/ost", logger(http.HandlerFunc(ostHandler)))

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("Failed to start server:", err)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path
	if fileName == "/" {
		fileName = "/index.html"
	} else {
		fileName = fileName + ".html"
	}

	filePath := filepath.Join("web", strings.TrimPrefix(fileName, "/"))

	if _, err := filepath.Glob(filePath); err != nil || len(filePath) == 0 {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}

func ostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("web", "/ost.docx"))
}

func logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"[%s] path: %s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
			r.RemoteAddr,
		)
	})
}
