package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {
	x = true
	fmt.Println(x)

	a := 10
	b := 20
	fmt.Println(a == b)

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)

	s := "small pr"
	fmt.Println(s)
	fmt.Println([]byte(s))

	fmt.Printf("%T\n", []byte(s))

	aa := []byte(s)
	for _, s := range aa {
		fmt.Println(s, string(s))
	}
	bb := "small pr"
	fmt.Printf("binary %b\n", bb)

	cc := "small pr"
	fmt.Printf("hexa %x\n", cc)

}
