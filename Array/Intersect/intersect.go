package main

import (
	"fmt"
	// "sort"
	// "sort"
)

func intersect(nums1 []int, nums2 []int) []int {

	result := make(map[int]bool)
	// index := 0
	array := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				array = append(array, nums1[i])
				

			} 
				
			
		}
	}
	// nums1 = [1,2,2,1], nums2 = [2,2]

	for _, num := range array {
		if !result[num] {
			result[num] = true
			// array = append(array, )
		}
	}
	
	// fmt.Println(result)
    array = []int{}
	for  num := range result {
		array = append(array, num)
	}
	return array

	// sort.Ints(nums1)
	// sort.Ints(nums2)
	// temp := 0
	// for i := 1; i < len(nums1); i++ {
	// 	if nums1[i] != nums1[temp] {
	// 		nums1[temp+1] = nums1[i]
	// 		temp++
	// 	}
	// }

	// array := make([]int , temp+1)

	// for i := 0; i < len(array); i++ {
	// 	array[i] = nums1[i]
	// }

	

	// temp = 0
	// for i := 1; i < len(nums2); i++ {
	// 	if nums2[i] != nums2[temp] {
	// 		nums2[temp+1] = nums2[i]
	// 		temp++
	// 	}
	// }

	// array2 := make([]int , temp+1)

	// for i := 0; i < len(array2); i++ {
	// 	array2[i] = nums2[i]
	// }
	// // fmt.Println(array2)


	// unique := make(map[int]bool)
	// result := []int{}

	// for i := 0; i < len(array); i++ {
	// 	for j := 0; j < len(array2); j++ {
	// 		if array[i] == array2[j] {
	// 			result = append(result, array[i]) 
	// 		}
				

			
				
	// 		}
	// 		for _, num := range result {
	// 			if !unique[num] { 
	// 				unique[num] = true
	// 				result = append(result, array[i]) }
	// 		}

	// 	}
	// 	keys := make([]int, 0, len(unique))
	// 	for key := range unique {
	// 		keys = append(keys, key)
	// 	}

	// return keys
	

	// return result


}

func main()  {

	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}

	result := intersect(nums1, nums2)
	fmt.Println(result)


}