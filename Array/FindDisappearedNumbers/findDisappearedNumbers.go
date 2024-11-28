package main

import (
	"fmt"
	// "sort"
)

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool)
    for _, v := range nums {
        m[v] = true
    }
	fmt.Println(m)

    res := make([]int, 0)
    for i := 1; i <= len(nums); i++ {
        if !m[i] {
            res = append(res, i)
			fmt.Println(i)
		
        }
    }
    return res
	// sort.Ints(nums)
		
	// array := []int{}
	// for i :=1; i <= len(nums); i++ {
	// 	array=append(array, i)
	// }
	// fmt.Println(array)
	// // index :=0
	// // arr := []int {}


	
	// elementMap := make(map[int]int)
	// result := []int{}

	// // Barcha elementlarni xaritada (map) hisoblang
	// for _, num := range nums {
	// 	elementMap[num]++
	// }
	// for _, num := range array {
	// 	elementMap[num]++
	// }

	// // Faqat bir marta uchragan elementlarni natijaga qo'shing
	// for key, count := range elementMap {
	// 	if count == 1 {
	// 		result = append(result, key)
	// 	}
	// }

	// return result
		
	}
	
// 1 2 3 4 7 8
func main() {
	nums := []int{1,2,2,2}
	resultt := findDisappearedNumbers(nums)
	fmt.Println(resultt)


}