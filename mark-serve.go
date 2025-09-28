package main

import (
	"flag"
	"github.com/gomarkdown/markdown"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ServerConfig struct {
	port             int
	markdownFilePath string
}

func main() {
	// --- 1. CONFIGURATION ---
	// Define and parse the command-line flags to configure the server.
	var serverConfig ServerConfig
	flag.IntVar(&serverConfig.port, "port", 8080, "The port for the server to listen on")
	flag.StringVar(&serverConfig.markdownFilePath, "file", "test.md", "The path to the markdown file to serve")
	flag.Parse()

	// --- 2. VALIDATION ---
	// Before starting, validate that the file exists and is accessible.
	// os.Stat is the most efficient way to do this.
	if _, err := os.Stat(serverConfig.markdownFilePath); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Error: The file '%s' was not found.", serverConfig.markdownFilePath)
		}
		log.Fatalf("Error checking file: %v", err)
	}

	// --- 3. ROUTING & HANDLERS ---
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This handler contains all the logic for a single request.
		log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)

		// a. Read the latest content from the file on disk.
		content, err := os.ReadFile(serverConfig.markdownFilePath)
		if err != nil {
			log.Printf("ERROR: could not read file %s: %v", serverConfig.markdownFilePath, err)
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		// b. Convert the markdown content to HTML.
		html := markdown.ToHTML(content, nil, nil)

		// c. Write the final HTML to the browser.
		// (Step 5 will replace this line with the template logic).
		io.ReadAll(template.HTML(co))

	})

	// --- 4. SERVER STARTUP ---
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.port),
		Handler: mux,
	}

	log.Printf("Starting server. Listening on http://localhost:%d", serverConfig.port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("FATAL: server failed to start: %v", err)
	}

	log.Println("Server stopped.")
}
