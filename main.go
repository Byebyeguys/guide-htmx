package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"guide/templates"
)

// Single channel for SSE - dead simple, no broker
var toastChan = make(chan string, 10)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/sse", handleSSE)
	http.HandleFunc("/trigger-toast", handleTriggerToast)
	http.HandleFunc("/delete-item", handleDeleteItem)
	http.HandleFunc("/form-submit", handleFormSubmit)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	templates.Index().Render(r.Context(), w)
}

func handleTriggerToast(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	toastType := r.FormValue("type")
	if toastType == "" {
		toastType = "success"
	}

	messages := map[string]string{
		"success": "Operation completed successfully!",
		"error":   "Something went wrong!",
		"info":    "Here's some information for you.",
		"warning": "Please be careful with this action.",
	}

	msg := messages[toastType]
	if msg == "" {
		msg = messages["info"]
	}

	toast := map[string]string{
		"type":    toastType,
		"message": msg,
	}
	data, _ := json.Marshal(toast)
	toastChan <- string(data)

	w.WriteHeader(http.StatusOK)
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mock 3 second delay to show loading state
	time.Sleep(3 * time.Second)

	toast := map[string]string{
		"type":    "error",
		"message": "Item deleted successfully!",
	}
	data, _ := json.Marshal(toast)
	toastChan <- string(data)

	w.WriteHeader(http.StatusOK)
}

func handleFormSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	// Simulate form processing
	toast := map[string]string{
		"type":    "success",
		"message": "Form submitted! Name: " + name + ", Email: " + email,
	}
	data, _ := json.Marshal(toast)
	toastChan <- string(data)

	// Return empty response - toast handles feedback
	w.WriteHeader(http.StatusOK)
}
