package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var r []int32
	fmt.Println(grades)
	for _, v := range grades {
		fmt.Println("for", v)
		if v < 40 {
			//grades[i] = v
			fmt.Println("if1", v)
		}
		if v > 40 {
			fmt.Println("if >", v)
			fmt.Println(v % 5)
			if v%5 > 3 {
				v = v + 5
				fmt.Println("if2", v)
			}
		}

		//fmt.Println("v", v)
	}
	return r

}

func main() {

	var grades = []int32{73, 67, 38, 33}
	result := gradingStudents(grades)
	fmt.Println(result)
}
