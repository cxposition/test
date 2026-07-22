package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Sum returns the sum of all integers in nums.
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Average returns the average of nums as float64.
// It returns 0 when nums is empty.
func Average(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}
	return float64(Sum(nums...)) / float64(len(nums))
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/healthx", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("okx"))
	})

	http.HandleFunc("/healthp", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("okx"))
	})

	http.HandleFunc("/healthpp", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("okx"))
	})

	http.HandleFunc("/healthppp", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("okx"))
	})

	http.HandleFunc("/scores", func(w http.ResponseWriter, _ *http.Request) {
		scores := []int{88, 95, 76, 90}
		resp := struct {
			Scores  []int   `json:"scores"`
			Sum     int     `json:"sum"`
			Average float64 `json:"average"`
		}{
			Scores:  scores,
			Sum:     Sum(scores...),
			Average: Average(scores...),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	})

	log.Println("server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
