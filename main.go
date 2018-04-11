package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		randomDuration := rand.Intn(3000)
		time.Sleep(time.Duration(randomDuration) * time.Millisecond)
		w.Write([]byte("pong"))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
