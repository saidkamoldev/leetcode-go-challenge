package main

import (
	"fmt"

)

func containsDuplicate(nums []int) bool  {

	// index := 0
	// count :=1
	// for i := 0; i < len(nums); i++ {
	// 	for j := i +1; j < len(nums); j++ {
	// 		if nums[j] == nums[i] {
	// 			count++
	// 		}
	// 	}
	// }
	// fmt.Println(count)
	// if count > 1 {
	// 	return true
	// } else {
	// 	return false
	// }
	
	seen := make(map[int]bool) // Xesh jadvali
    for _, num := range nums {
        if seen[num] { 
            return true
        }
        seen[num] = true 
    }
    return false

}

func main() {
	nums := []int{2,14,18,22}
	// 1,2,3,1
	// 1,2,3,4

    booool :=containsDuplicate(nums)
	fmt.Println(booool)



}