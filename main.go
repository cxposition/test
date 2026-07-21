package main

import "fmt"

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
	scores := []int{88, 95, 76, 90}
	fmt.Printf("scores=%v\n", scores)
	fmt.Printf("sum=%d\n", Sum(scores...))
	fmt.Printf("avg=%.2f\n", Average(scores...))
}
