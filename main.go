package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metrics/middleware.go
var (
	httpRequestsResponseTime prometheus.Summary
)

func init() {
	httpRequestsResponseTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "http",
		Name:      "response_time_seconds",
		Help:      "Request response times",
	})

	prometheus.MustRegister(httpRequestsResponseTime)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		httpRequestsResponseTime.Observe(float64(time.Since(start).Seconds()))
	})
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	randomDuration := rand.Intn(50)
	time.Sleep(time.Duration(randomDuration) * time.Millisecond)
	w.Write([]byte("pong"))
}

func main() {

	http.Handle("/ping", Middleware(http.HandlerFunc(HandlePing)))

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8081", nil))

}
