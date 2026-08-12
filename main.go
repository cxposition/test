package main

import (
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

	http.HandleFunc("/healthzzzz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	log.Println("server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
