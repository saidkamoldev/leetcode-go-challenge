package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {

		// i,j := 0, 0
		// last := len(students)-1
		// count:=0
		

		// for i <len(students) && j < len(sandwiches) {
		// 	if students[i] == sandwiches[j] {
		// 		i++
		// 		j++
		// 		count++
		// 		continue
				
		// 	}

			

		// 	if students[i] != sandwiches[j] {
		// 		temp := students[i]
		// 		students[i] = students[last]
		// 		students[last] = temp
		// 	}
		// 	fmt.Println(students)
		// 	if students[i] == students[last] {
		// 		break
		// 	}
		// }
		// // count = len(students) - count

		// return count 
		i, j := 0, 0 
		rejected := 0 
	
		for i < len(students) {
			if students[i] == sandwiches[j] {
				
				students = append(students[:i], students[i+1:]...)
				j++ 
				rejected = 0 
			} else {
				// Talaba oxiriga o'tadi
				students = append(students[1:], students[0])
				rejected++
			}
			fmt.Println(students)
			
			if rejected == len(students) {
				break
			}
		}
	
		return len(students) 

}

func main()  {
	students  := []int{1,1,0,0}
	sandwiches := []int{0,1,0,1}

	result := countStudents(students, sandwiches)
	fmt.Println(result)
}