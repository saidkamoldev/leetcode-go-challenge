package main

import "fmt"

func twoSum(nums []int, target int) []int {

	result := []int{}
OuterLoop:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ { // j := i + 1 qilib yozamiz, o'zini o'zi tekshirmaslik uchun
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				break OuterLoop
			}
		}
	}

	return result
}

// }

func main() {
	var n, target int

	fmt.Println("Elementlar sonini kiriting:")
	fmt.Scan(&n) // elementlar sonini olish

	nums := make([]int, n)

	fmt.Println("Massiv elementlarini kiriting:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i]) // massivni kiritish
	}

	fmt.Println("Targetni kiriting:")
	fmt.Scan(&target) // targetni kiritish

	result := twoSum(result)

}
