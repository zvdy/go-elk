package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func generateRandomLog() map[string]string {
	timestamp := time.Now().Format(time.RFC3339)
	level := []string{"INFO", "WARN", "ERROR"}[rand.Intn(3)]
	message := fmt.Sprintf("This is a %s message", level)
	return map[string]string{
		"timestamp": timestamp,
		"level":     level,
		"message":   message,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	elasticURL := "http://elasticsearch:9200/random_logs/_doc"

	// Retry mechanism to wait for Elasticsearch to be ready
	for {
		resp, err := http.Get("http://elasticsearch:9200")
		if err == nil && resp.StatusCode == 200 {
			resp.Body.Close()
			break
		}
		fmt.Println("Waiting for Elasticsearch to be ready...")
		time.Sleep(5 * time.Second)
	}

	for {
		log := generateRandomLog()
		logJSON, err := json.Marshal(log)
		if err != nil {
			fmt.Println("Error marshaling log to JSON:", err)
			continue
		}

		req, err := http.NewRequest("POST", elasticURL, bytes.NewBuffer(logJSON))
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending log to Elasticsearch:", err)
			continue
		}
		resp.Body.Close()

		time.Sleep(1 * time.Second) // Optional: Add a delay to control the log generation rate
	}
}
