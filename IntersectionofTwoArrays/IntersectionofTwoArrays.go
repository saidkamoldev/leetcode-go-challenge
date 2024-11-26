package main

import (
	"fmt"

)

func intersection(nums1 []int, nums2 []int) []int {
	unique := make(map[int]bool)
	array := []int{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				array = append(array, nums1[i]) 
			}
				

			
				
			}
			for _, num := range array {
				if !unique[num] { 
					unique[num] = true
					array = append(array, nums1[i]) }
			}

		}
		keys := make([]int, 0, len(unique))
		for key := range unique {
			keys = append(keys, key)
		}

	return keys



}
	

func main() {
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	resultt := intersection(nums1, nums2)
	fmt.Println(resultt)


}