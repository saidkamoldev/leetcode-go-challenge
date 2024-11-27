package main

import (
	"sort"
	"fmt"
)

func thirdMax(nums []int) int {
	
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	elemCounted := 1
	prevElem := nums[0]

	for i := 1; i < len(nums); i++ {
	
		if nums[i] != prevElem {
			elemCounted++
			prevElem = nums[i]
		}


		if elemCounted == 3 {
			return nums[i]
		}
	}

	
	return nums[0]
	// for i := 0; i < len(nums); i++ {

	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] < nums[j] {
	// 			temp := nums[j]
	// 			nums[j] = nums[i]
	// 			nums[i] = temp
	// 		}
	// 	}

	// }

	
	
	// temp := 0
	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] != nums[temp] {
	// 		nums[temp+1] = nums[i]
	// 		temp++
	// 	}
	// }

	// array := make([]int , temp+1)

	// for i := 0; i < len(array); i++ {
	// 	array[i] = nums[i]
	// }

	// result := array[0]
    // if len(array)<= 2 {
	// 	for i := 0; i < len(array); i++ {
	// 		if array[i] > result{
	// 			result = array[i]
				
	// 		}
	// 	}
	// 	return result
	// }

	// answer:= array[2]
	
	// return answer
}

func main() {
	nums := []int {1,1,2}
	resultt := thirdMax(nums)
	fmt.Println(resultt)


}