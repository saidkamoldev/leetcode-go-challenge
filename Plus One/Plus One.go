package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i < 0; i-- {
		digits[i]++

		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {

	nums := []int{1, 2, 3, 4, 5} // Original slice
	// // val := 2                           // O'chirilishi kerak bo'lgan qiymat

	newLength := plusOne(nums) // removeElement funksiyasini chaqirish

	fmt.Println(newLength)
}
