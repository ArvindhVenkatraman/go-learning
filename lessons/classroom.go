package main

import "fmt"

var y = "hi"
var z = "Hi Arvindh"
var a = `Something there "okay i got it" then what happens @#$%^&*()`

type mytype int

var myvar mytype
var myvar2 string

func main() {
	n, _ := fmt.Println("hello world")
	fmt.Println(n)
	x := 45
	fmt.Println(x)
	x = 27
	fmt.Println(x)
	fmt.Println(y)
	//z := 10
	//_, err := foo(z)
	//z = 10
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println("Message : ", a)
	fmt.Printf("Raw String Type: %T\n", a)

	juju := "Arjun"
	harshy := "Harshad"
	poppy := "Rama"

	fmt.Printf("Juju hexa - %X\t%#X\t%b\n", juju, juju, juju)
	fmt.Printf("Harshy hexa - %X\t%#X\t%b\n", harshy, harshy, harshy)
	fmt.Printf("Poppy hexa - %X\t%#X\t%b\n", poppy, poppy, poppy)

	myvar2 = string(myvar)
	fmt.Println("This is myvar2", myvar2)

	foo()

}

func foo() {
	n, _ := fmt.Println("I'm in foo")
	fmt.Println(n)
	fmt.Println(y)
	//z = 11
	fmt.Println(z)

}

// Class room notes
//ctrl + shift + R to run the code
// := short declaration. x := 42 (declares the type and value)
// x = 55 (once declared can be used with = )
// operand and operator
// operator is :=
// operand is x and 42 ( example x := 42)
// var is generally used for variable declaration outside of function. like global variable.
// if you want to use instead function, use short declaration.
// raw string literal `` pretty much outputs everything declared in that
