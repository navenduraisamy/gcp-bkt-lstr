package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/healthz", healthCheckHandler)
	mux.HandleFunc("/api/gcp/buckets", listBucketsHandler)

	fmt.Printf("starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID != "" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func listBucketsHandler(w http.ResponseWriter, r *http.Request) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		http.Error(w, "Project ID empty", http.StatusInternalServerError)
		return
	}

	buckets, err := listBuckets(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	responseJson, _ := json.Marshal(buckets)
	fmt.Fprintf(w, string(responseJson))
}

func listBuckets(projectID string) ([]string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var buckets []string
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
		fmt.Printf("Bucket: %v\n", battrs.Name)
	}
	return buckets, nil
}
