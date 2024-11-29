package main

import (
	"fmt"

)


func matrixReshape(mat [][]int, r int, c int) [][]int {
	
	   // 1. Asl matritsaning oʻlchamini aniqlaymiz
	   m, n := len(mat), len(mat[0]) // `m` qatorlar soni, `n` ustunlar soni
	fmt.Println(m)
	fmt.Println(n)
	   // 2. Agar elementlar soni mos kelmasa, asl matritsani qaytaramiz
	   if m*n != r*c {
		   return mat
	   }
   
	   // 3. Yangi matritsani e'lon qilish
	   reshapedMatrix := make([][]int, r)
	   for i := range reshapedMatrix {
		   reshapedMatrix[i] = make([]int, c)
	   }
   
	   // 4. Elementlarni yangi matritsaga koʻchirish
	   for i := 0; i < m*n; i++ {
		   reshapedMatrix[i/c][i%c] = mat[i/n][i%n]
	   }
   
	   // 5. Yangi matritsani qaytarish
	   return reshapedMatrix
}

func main()  {
	mat := [][]int{
        {1, 2},
        {3, 4},
    }

    // Qatorlar va ustunlar soni
    r, c := 1, 4
	result := matrixReshape(mat, r,c)
	fmt.Println(result)
}