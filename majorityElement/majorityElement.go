package main

import "fmt"

func majorityElement(nums []int) int {

	candidate, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate

}

func main() {
	nums := []int{3, 2, 3}
	newLength := majorityElement(nums)
	fmt.Println(newLength)

}

// count := 0
// index := 0

// for i := 0; i < len(nums); i++ {
// 	temp := 0
// 	for j := 0; j < len(nums); j++ {
// 		if nums[i] == nums[j] {
// 			temp++
// 		}
// 	}
// 	if temp > count {
// 		count = temp
// 		index = i
// 	}
// }

// return nums[index]
