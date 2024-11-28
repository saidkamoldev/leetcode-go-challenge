package main

import (
	"fmt"
	"sort"
)


func arrayPairSum(nums []int) int {
    sort.Ints(nums)

	sum := 0
	index :=0

	for i := 1; i < len(nums); i++ {
		if nums[index] >= nums[i] {
			sum += nums[index]
		} else if  nums[index] <= nums[i]{
			sum += nums[index]
		}
		index+=2
		i++
	}
return sum
}

func main()  {
	nums := []int{6,2,6,5,1,2}
	result := arrayPairSum(nums)
	fmt.Println(result)
}