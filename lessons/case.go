package main

import "fmt"

func main() {
	// missing switch expression is always considered as true.
	switch false {
	case false:
		fmt.Println("This is true")
		//fallthrough
		// So the case statement only considers one of the true case.
		//If the true case is available then it considers the case
		// you can use fallthrough
	case (27 == 28):
		fmt.Println("This is 27" +
			"Omgodzilla")

	default:
		fmt.Println("Im in default")
	}
}
