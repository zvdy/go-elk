package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func generateRandomLog(r *rand.Rand) map[string]interface{} {
	timestamp := time.Now().Format(time.RFC3339)
	clientIP := fmt.Sprintf("192.168.%d.%d", r.Intn(256), r.Intn(256))
	requestURI := fmt.Sprintf("/api/v1/resource/%d", r.Intn(100))
	responseStatus := []int{200, 201, 400, 401, 403, 404, 500}[r.Intn(7)]
	responseTime := r.Intn(1000) // in milliseconds
	apiProduct := []string{"ProductA", "ProductB", "ProductC"}[r.Intn(3)]
	org := []string{"Org1", "Org2", "Org3"}[r.Intn(3)]
	env := []string{"Dev", "Staging", "Prod"}[r.Intn(3)]

	log := map[string]interface{}{
		"timestamp":       timestamp,
		"client_ip":       clientIP,
		"request_uri":     requestURI,
		"response_status": responseStatus,
		"response_time":   responseTime,
		"api_product":     apiProduct,
		"org":             org,
		"env":             env,
	}

	if responseStatus != 200 && responseStatus != 201 {
		faultCode := []string{"Fault1", "Fault2", "Fault3"}[r.Intn(3)]
		faultPolicy := []string{"Policy1", "Policy2", "Policy3"}[r.Intn(3)]
		log["fault_code"] = faultCode
		log["fault_policy"] = faultPolicy
	}

	return log
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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
		log := generateRandomLog(r)
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
