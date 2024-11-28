package main

import "fmt"

func fib(n int) int {
	f1 := 0 
	f2 := 1 
    if n == 0{
        return f1
    }
// 	// for i := 0; i < n-1; i++ {
//     leng :=0
// 	for leng < n-1 {
// 		temp := f1 + f2 
// 		// fmt.Println(temp)
// 		f1 = f2     
// 		// fmt.Println(f1)   
// 		f2 = temp  
// 		// fmt.Println(f2)  
//         leng++
//         }
// 	// }

// 	return f2

    leng :=1
	for leng < n {
	
		temp := f1 + f2 
	
		f1 = f2     
	
		f2 = temp  
		   leng++
	}

	return f2
}

func main()  {
	n := 3
	result := fib(n)
	fmt.Println(result)
}