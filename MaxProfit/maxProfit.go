package main

import (
	"fmt"

)

func maxProfit(prices []int) int {

	buy :=prices[0]
	sell :=0

	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else if prices[i] - buy > sell  {
			sell = prices[i]-buy
		}
	}
	fmt.Println(sell)
	// result := 0
	// for i := 1; i < len(prices); i++ {

	// 	if prices[i] < prices[result]{
	// 		result = i
			
	// 	}
	// }
	// // fmt.Println(result)
	// index :=0
	// for i := result; i < len(prices); i++ {
	// 	if result < prices[i] {
	// 		index = i
	// 	} 
	// }
	// // fmt.Println(index)
	// if index <= result {
	// 	result = 0
	// } else {
	// 	result = index
	// }


	return sell
}
	

func main() {
	prices := []int{3,3}
	resultt := maxProfit(prices)
	fmt.Println(resultt)


}