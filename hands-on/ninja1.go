package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

type mytype int

var myvar mytype

var myvarY int

func main() {
	execrise1()
	execrise2()
	execrise3()
	execrise4()
	execrise5()

}

func execrise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("----------- Execrise 1 -------------------")
	fmt.Println("Execrise 1 Printing single line :", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("------------------------------------------")
}

func execrise2() {
	fmt.Println("----------- Execrise 2 -------------------")
	fmt.Println("Execrise 2 Printing single line :", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("------------------------------------------")
}

func execrise3() {
	fmt.Println("------- Execrise 3 -----------")
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	fmt.Println("------------------------------")
}

func execrise4() {
	fmt.Println("------- Execrise 4 -----------")
	fmt.Println("Value", myvar)
	fmt.Printf("%T\n", myvar)
	myvar = 27
	fmt.Println("Value", myvar)
	fmt.Println("------------------------------")
}

func execrise5() {
	fmt.Println("------- Execrise 5 -----------")
	fmt.Println("Value", myvar)
	fmt.Printf("%T\n", myvar)
	myvar = 27
	fmt.Println("Value", myvar)
	myvarY = int(myvar)
	fmt.Println(myvarY)
	fmt.Println("------------------------------")
}
