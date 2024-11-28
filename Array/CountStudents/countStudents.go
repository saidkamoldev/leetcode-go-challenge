package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {

	count :=0
	index :=0
	index2 :=0
	for index < len(students){

		if students[index2] == 1 && sandwiches[index] == 1  {
			count++
		
		} else if  sandwiches[index] == 0{
			index2++
	
		}
	
	
		if index2 == len(students) {
			break
		}
		index++
	}
return count  

}

func main()  {
	students  := []int{1,1,1,0,0,1}
	sandwiches := []int{1,0,0,0,1,1}

	result := countStudents(students, sandwiches)
	fmt.Println(result)
}