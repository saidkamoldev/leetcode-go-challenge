package main

import (
	"fmt"
	"sort"
)


func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
    contentChildren := 0
    cookieIndex := 0

    for cookieIndex < len(s) && contentChildren < len(g) {
        if s[cookieIndex] >= g[contentChildren] {
            contentChildren++
        }
        cookieIndex++
    }

   return contentChildren
}


func main()  {
	g := []int{1,2,3}
	s := []int{1,1}

	result := findContentChildren(g,s)
	fmt.Println(result)
}