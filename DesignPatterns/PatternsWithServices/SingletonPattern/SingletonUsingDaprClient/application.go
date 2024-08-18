package main

import (
	"encoding/json"
	"log"
	"net/http"

	// "singletonObj/publisher"

	"github.com/gorilla/mux"
)

func main() {
	r := setUpRouter()
	http.ListenAndServe(":8080", r)
}

func setUpRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/subscribe", subRoute)
	return r
}

func subRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data string

	err := json.NewDecoder(r.Body).Decode(&data)
	log.Println("Received Message: ", data)

	if err != nil {
		log.Printf("Error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Access the singleton Dapr client
	// publisher.GetInstance().client

	// Use the client as needed
	// Example: client.PublishEvent(ctx, pubsubName, topicName, data)
}
