package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			target = i
			return target
		}

		if nums[i] > target {
			return i
		}
	}

	return len(nums)
}

func main() {
	nums := []int{1, 2, 3, 4}
	target := 5

	newLength := searchInsert(nums, target)
	fmt.Println("Qolgan elementlar soni:", newLength)

}
