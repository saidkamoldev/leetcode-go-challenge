package main

import "fmt"


func findMaxConsecutiveOnes(nums []int) int {
	
	count := 0
	i:=0
	temp :=0

   for i < len(nums) {
	 
	 if nums[i] == 1  {
		count++
	 } else if nums[i] ==0 {
		count = 0
	 }	 
	 if temp<count {
		temp=count
		// fmt.Println(temp)
	 }

	 i++

   }

	return temp


}

func main()  {
	nums := []int{1,0,1,1,0,1}
	result := findMaxConsecutiveOnes(nums)
	fmt.Println(result)
}