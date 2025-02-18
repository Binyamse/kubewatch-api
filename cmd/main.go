package main

import (
	"fmt"
	"net/http"

	"kubewatch-api/pkg/monitor"
)

func main() {
	http.Handle("/metrics", http.HandlerFunc(monitor.MetricsHandler))
	fmt.Println("Starting Kubernetes API Monitor on :8080")
	http.ListenAndServe(":8080", nil)
}
