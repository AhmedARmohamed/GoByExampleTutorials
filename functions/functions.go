package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

// Varidiac functions

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, nums := range nums {
		total += nums
	}
	fmt.Println(total)
}

func main() {
	// res := plus(1, 3)
	// fmt.Println("1 + 3 = ", res)

	sum(1,23, 5, 6, 7, 9, 0 ,8, 6)
}
