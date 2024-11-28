package main

import (
	"fmt"

)

func Merge(nums1 []int, m int, nums2 []int, n int)  {

	index :=m

	for i := 0; i < len(nums2); i++ {
	
			nums1[index] = nums2[i]
			index++

	}
	for i := 0; i < len(nums1); i++ {
		for j := i+1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				temp := nums1[j]
				nums1[j] = nums1[i]
				nums1[i] =temp
			}
		}
	}

}

func main() {
	nums1 := []int{}
	nums2 := []int{1}
	m:=0
	n := 1
    Merge(nums1, m, nums2, n)


}