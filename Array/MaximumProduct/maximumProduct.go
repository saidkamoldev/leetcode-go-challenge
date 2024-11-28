package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
    
    array := []int{}
    
    for i := 0; i < 3; i++ {
        array = append(array, nums[i])
    }
    
    for i := len(nums) - 1; i > len(nums) - 1 - 3 && i >= 3; i-- {
        array = append(array, nums[i])
    }
    fmt.Println(array)
    result := nums[0] * nums[1] * nums[2]
    for i := 0; i < len(array); i++ {
        for j := i + 1; j < len(array); j++ {
            for k := j + 1; k < len(array); k++ {
                result = max(result, array[i] * array[j] * array[k])
            }
        } 
    }
    return result

}

func main()  {
	num := []int{1,2,3,4}
	result := maximumProduct(num)
	fmt.Println(result)
}