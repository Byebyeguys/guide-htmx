package main

import (
	"fmt"
	"net/http"
)

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Required headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	// Send initial connection event
	fmt.Fprintf(w, "event: connected\ndata: {\"status\":\"connected\"}\n\n")
	flusher.Flush()

	for {
		select {
		case <-r.Context().Done():
			// Client disconnected
			return
		case msg := <-toastChan:
			// SSE format: "event: name\ndata: payload\n\n"
			fmt.Fprintf(w, "event: sse-toast\ndata: %s\n\n", msg)
			flusher.Flush()
		}
	}
}
