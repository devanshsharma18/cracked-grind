package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ServerConfig struct {
	port             int
	markdownFilePath string
}

func main() {
	// TODO (Step 2): Use the `flag` package to parse command-line arguments.
	// We need a flag for the port number (e.g., `-port=8080`)
	// and a flag for the markdown file path (e.g., `-file=README.md`).
	var serverConfig ServerConfig
	flag.IntVar(&serverConfig.port, "port", 8080, "Server port")
	flag.StringVar(&serverConfig.markdownFilePath, "file", "test.md", "File name")
	flag.Parse()
	// A basic check to see if the file exists.
	if _, err := os.Open(serverConfig.markdownFilePath); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist.", "markdownFilePath")
	}
	defer file.Close()
	// Create a new HTTP server multiplexer (router).
	mux := http.NewServeMux()

	// This is the main handler for rendering the markdown file.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)
		// TODO (Step 3): Read the content of `markdownFilePath`.
		// Handle potential errors.
		scanner := bufio.NewReader(serverConfig.markdownFilePath)
		os.OpenFile(serverConfig.markdownFilePath)
		// TODO (Step 4): Convert the markdown content to HTML.
		// This is where you will use the `gomarkdown/markdown` library.
		// Research how to use the `mdtopdf.MarkdownToHTML` function.

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
		Addr:    ":" + port,
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

	log.Printf("Starting server on http://localhost:%s", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}

	// TODO (Step 8): Implement Graceful Shutdown.
	// 1. Create a channel to listen for `os.Interrupt` signal.
	// 2. When the signal is received, call `srv.Shutdown()` with a context.
	// 3. Wait for the server to shut down cleanly.

	log.Println("Server stopped.")
}
