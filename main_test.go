package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(1, 2, 3, 4)
	if got != 10 {
		t.Fatalf("Sum() = %d, want 10", got)
	}
}

func TestAverage(t *testing.T) {
	got := Average(2, 4, 6, 8)
	if got != 5 {
		t.Fatalf("Average() = %v, want 5", got)
	}
}

func TestAverageEmpty(t *testing.T) {
	got := Average()
	if got != 0 {
		t.Fatalf("Average() = %v, want 0", got)
	}
}
