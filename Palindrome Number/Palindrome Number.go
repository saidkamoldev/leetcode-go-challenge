package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	last := x
	sum := 0
	num := 0
	var b bool
	for x != 0 {
		sum = x % 10
		num = num*10 + sum
		x = x / 10
	}

	if num > 0 && num == last {
		b = true
	} else {
		b = false
	}
	return b
}

func main() {
	number := 0 // bu yerda tekshirmoqchi bo'lgan sonni kiriting
	if isPalindrome(number) {
		fmt.Printf("%d - bu palindrom\n", number)
	} else {
		fmt.Printf("%d - bu palindrom emas\n", number)
	}
}
