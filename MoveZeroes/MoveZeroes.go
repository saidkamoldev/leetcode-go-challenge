package main

import (
	"fmt"

)

func moveZeroes(nums []int)  {

	index :=0
	for i := 1; i < len(nums); i++ {
		if nums[index] != 0 && (nums[i] !=0 || nums[i] ==0) {
			index++
		} else if nums[index] == 0 && nums[i] !=0 {
			nums[index]= nums[i]
			nums[i]=0
			index++

		}
	}

	fmt.Println(nums)
	// array := make([]int, len(nums))
	// index:=0
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != 0 {
	// 		array[index] = nums[i]
	// 		index++
	// 	} 
	// }

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 0 {
	// 		array[index] = nums[i]
	// 		index++
	// 	}
        
	// }
	// for i := 0; i < len(nums); i++ {
	// 	nums[i]=array[i]
	// }
	

}

func main() {
	nums := []int{1,0,1}


    moveZeroes(nums)


}