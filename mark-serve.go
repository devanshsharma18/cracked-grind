package main

import (
	"flag"
	"fmt"
	"github.com/gomarkdown/markdown"
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

	//Define and parse the command-line flags required to configure the server on startup
	//This allows user to specify the port number and markdown file to the serve.
	var serverConfig ServerConfig
	flag.IntVar(&serverConfig.port, "port", 8080, "Server port")
	flag.StringVar(&serverConfig.markdownFilePath, "file", "test.md", "File name")
	flag.Parse()

	//Before starting the server, we check if the target file exists and can be accessed.
	// In case the file does not exist or cannot be opened, we exit with errors
	file, err := os.Open(serverConfig.markdownFilePath)
	if err != nil {
		if os.IsNotExist(err) {

			// If the file does not exist, log a "Does not exist" message.
			log.Fatalf("File '%s' does not exist.", serverConfig.markdownFilePath)
		}
		//Handle all other potential errors (e.g., permission denied).
		log.Fatalf("An unexpected error occurred while opening file: %v", err)
	}
	//Deferring the file closure to the final steps of the program
	defer file.Close()

	// Create a new HTTP server multiplexer (router).
	mux := http.NewServeMux()

	// This is the main handler for rendering the markdown file.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)

		//Reading the contents of the file submitted by the user.
		// Handling potential errors where the file is not accessible.
		content, err := os.ReadFile(serverConfig.markdownFilePath)

		if err != nil {
			log.Printf("ERROR: could not read file %s: %v", serverConfig.markdownFilePath, err)
			// Sending a user an error message on the browser level.
			http.Error(w, "500 Internal Server Error: Could not read file.", http.StatusInternalServerError)
			return
		}

		// Convert the raw markdown bytes into HTML bytes using the markdown library's default settings.
		html := markdown.ToHTML(content, nil, nil)
		w.Write(html)

		// TODO (Step 5): Inject the generated HTML into our `template.html`.
		// You will need to read `template.html` and replace a placeholder
		// (e.g., `<!--CONTENT-->`) with your generated HTML.
		// Then, write the final HTML to the `http.ResponseWriter`.
		// For now, let's just write a simple message.
		fmt.Fprint(w, "Hello, World!")
	})

	// TODO (Step 6): Create a WebSocket handler at `/ws`.
	// This will involve:
	// 1. "Upgrading" the HTTP connection to a WebSocket connection.
	// 2. Storing the connection so we can send messages to it later.
	// 3. Handling connection closures.
	// This is a major step, so break it down!

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.port),
		Handler: mux,
	}

	// TODO (Step 7): Start a goroutine to watch the markdown file for changes.
	// This will be our file watcher from scratch (the stretch goal).
	// 1. Create a channel to signal changes.
	// 2. In a new goroutine, start a loop that periodically checks the file's
	//    modification time using `os.Stat`.
	// 3. If the modification time has changed, send a signal on the channel.
	// 4. When the signal is received, send a "reload" message to all connected
	//    WebSocket clients.

	log.Printf("Starting server on http://localhost:%s", strconv.Itoa(serverConfig.port))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", strconv.Itoa(serverConfig.port), err)
	}

	// TODO (Step 8): Implement Graceful Shutdown.
	// 1. Create a channel to listen for `os.Interrupt` signal.
	// 2. When the signal is received, call `srv.Shutdown()` with a context.
	// 3. Wait for the server to shut down cleanly.

	log.Println("Server stopped.")
}
