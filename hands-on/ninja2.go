package main

import "fmt"

func main() {
	_execrise2()
	_execrise4()
	_execrise5()
	_execrise6()
}

func _execrise2() {

	a := (27 == 27)
	b := (27 <= 27)
	c := (27 >= 27)
	d := (27 != 27)
	e := (27 < 27)
	f := (27 > 27)
	println(a, b, c, d, e, f)
}

func _execrise4() {
	a := 42
	fmt.Printf("%d\t\t%b\t\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", b, b, b)

}

func _execrise5() {
	a := `"I'm printing a raw string 
literal so that you know what im talking about
thank you'"`
	fmt.Println(a)

}

func _execrise6() {
	println(a, b, c, d)

}
