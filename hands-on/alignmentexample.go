package main

import (
	"fmt"
	"strings"
)

func rightjust(s string, n int, fill string) string {
	return strings.Repeat(fill, n) + s
}

func leftjust(s string, n int, fill string) string {
	return s + strings.Repeat(fill, n)
}

func center(s string, n int, fill string) string {
	div := n / 2

	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
}

func main() {
	str := "Hello World!"
	fmt.Println("Original : ", str)

	fmt.Println("Adjust 20 spaces right : [" + rightjust(str, 20, " ") + "]")

	fmt.Println("Adjust 20 spaces left : [" + leftjust(str, 20, " ") + "]")

	fmt.Println("Center 10 spaces : [" + center(str, 20, " ") + "]")

	fmt.Println("Center * fill : [" + center(str, 10, "*") + "]")
}
