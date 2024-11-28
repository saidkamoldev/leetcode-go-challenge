package main

import (
	"fmt"

)

func missingNumber(nums []int) int {

	index :=0
	count :=0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		index += (i + 1)
	}

	return index - count
}
	

func main() {
	nums := []int{9,6,4,2,3,5,7,0,1}
	resultt := missingNumber(nums)
	fmt.Println(resultt)


}