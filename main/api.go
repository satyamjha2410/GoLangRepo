package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/couchbase/gocb/v2"
)

// Couchbase configuration
const (
	couchbaseEndpoint = "couchbase://localhost"
	bucketName        = "your_bucket_name"
	username          = "your_username"
	password          = "your_password"
)

// Record struct to represent a data record
type Record struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

func fetchRecordHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID parameter from the URL
	vars := mux.Vars(r)
	recordID := vars["id"]

	// Connect to Couchbase
	cluster, err := gocb.Connect(couchbaseEndpoint, gocb.ClusterOptions{
		Username: username,
		Password: password,
	})
	if err != nil {
		http.Error(w, "Error connecting to Couchbase", http.StatusInternalServerError)
		return
	}
	defer cluster.Close(nil)

	// Open the bucket
	bucket := cluster.Bucket(bucketName)

	// Get the document from Couchbase
	var record Record
	_, err = bucket.Get(recordID, &record)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching record with ID %s: %v", recordID, err), http.StatusNotFound)
		return
	}

	// Convert the record to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the route for fetching a record by ID
	router.HandleFunc("/try/go/fetch/{id}", fetchRecordHandler).Methods("GET")

	// Start the server
	port := "8080"
	log.Printf("Server listening on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
