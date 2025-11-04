package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	digits := []int{1, 2, 3}
	digits = plusOne(digits)
	fmt.Printf("%v\n", digits)

	digits = []int{4, 3, 2, 1}
	digits = plusOne(digits)
	fmt.Printf("%v\n", digits)

	digits = []int{9}
	digits = plusOne(digits)
	fmt.Printf("%v\n", digits)
}
