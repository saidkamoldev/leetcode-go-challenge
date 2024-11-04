package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	size := len(digits)

	for i := size - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}

	newDigits := make([]int, size+1)
	newDigits[0] = 1

	return newDigits
}

func main() {
	nums := []int{1, 2, 3, 4}
	newLength := plusOne(nums)

	fmt.Println(newLength) // Output: [4, 3, 2, 2]
}
