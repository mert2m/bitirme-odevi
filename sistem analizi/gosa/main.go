package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	})

	responseTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "http_response_time_seconds",
		Help: "HTTP response time in seconds",
	})
)

func main() {
	// Prometheus metriklerini kaydetmek için kaydediciye register ediyoruz
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(responseTime)

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/staj", handleStaj)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc()

	startTime := time.Now()

	fmt.Fprint(w, "Merhaba Go!")

	elapsedTime := time.Since(startTime).Seconds()
	responseTime.Observe(elapsedTime)
}

func handleStaj(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc()

	startTime := time.Now()

	rand.Seed(time.Now().UnixNano())

	countries := []string{
		"Turkey",
		"United States",
		"Germany",
		"France",
		"United Kingdom",
		"Japan",
		"Brazil",
		"Russia",
		"Australia",
		"Canada",
		"Italy",
		"China",
	}

	// Rastgele bir ülke seçimi
	randomIndex := rand.Intn(len(countries))
	selectedCountry := countries[randomIndex]

	// JSON dönüşü
	jsonData, err := json.Marshal(selectedCountry)
	if err != nil {
		fmt.Println("JSON marshaling error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// HTTP yanıtı
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	elapsedTime := time.Since(startTime).Seconds()
	responseTime.Observe(elapsedTime)
}
