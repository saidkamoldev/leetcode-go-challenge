package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
		}
	}

	return result
}

func main() {
	nums := []int{3, 2, 2, 3, 4, 2, 5} // Original slice
	val := 2                           // O'chirilishi kerak bo'lgan qiymat

	newLength := removeElement(nums, val) // removeElement funksiyasini chaqirish
	fmt.Println("Qolgan elementlar soni:", newLength)
	fmt.Println("Yangi slice:", nums[:newLength]) // Qolgan elementlar
}
