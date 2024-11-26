package main

import (
	"fmt"

)

func distributeCandies(candyType []int) int {

	count := 0

	unique := make(map[int]bool)

	index :=0

		// if candyType[index] == candyType[i] {
			
		// }
		for _, num := range candyType {
			if !unique[num] { 
				unique[num] = true
				count++
				if count == len(candyType) / 2 {
					index = len(unique)
					return index
				}
				// candyType = append(candyType, candyType[i]) 
			}
		}

	return count
}
	

func main() {
	candyType := []int{1,1,2,3}

	result := distributeCandies(candyType)
	fmt.Println(result)


}