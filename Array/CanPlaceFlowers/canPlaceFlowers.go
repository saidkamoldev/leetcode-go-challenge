package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count :=0
	index :=0
	a := 2
	b := 3
	temp:=0
	// c := 2

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			temp++
		}
	}
	// fmt.Println(count)
	// fmt.Println(n*2)
	if n * 2 + 1 > temp {
		return false
	} 
	
     
	for i := 1; i < len(flowerbed); i++ {
		if index == 0 && flowerbed[index]==0 && i == 1 && flowerbed[i] == 0 {
			count++
		} else if flowerbed[i] ==0&& flowerbed[a] == 0 && flowerbed[b] == 0 {
			count++
		}


		if b > len(flowerbed) && a > len(flowerbed) {
			b++
			a++
		}
		

	}
	// fmt.Println(count)
	if count >= n {
		return true
	} else {
		return false
	}
	// return count



}
func main() {
	flowerbed := []int {0,0,1,0,0}
	n := 2
	resultt := canPlaceFlowers(flowerbed, n)
	fmt.Println(resultt)


}