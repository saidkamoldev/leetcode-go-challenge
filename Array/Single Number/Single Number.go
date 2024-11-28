package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num // XOR operatori
	}
	return result

}

func main() {
	var n int

	fmt.Println("Elementlar sonini kiriting:")
	fmt.Scan(&n) // elementlar sonini olish

	nums := make([]int, n)

	fmt.Println("Massiv elementlarini kiriting:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i]) // massivni kiritish
	}

	sum := singleNumber(nums)
	fmt.Println(sum)

}
