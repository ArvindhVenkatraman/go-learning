package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb
	gb
)

func main() {
	fmt.Printf("%b\t\t%b\t\t%b\n", kb, mb, gb)

}
