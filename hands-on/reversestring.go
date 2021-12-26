package main

import (
	"fmt"
	"strings"
)

var mystr []string //slice of string

func main() {
	str := "34o3ll435e  39h   "  //Your input string
	str = strings.TrimSpace(str) //trims the spaces
	for _, v := range str {
		if v > 65 {
			mystr = append(mystr, string(v)) //for every alphabet append the character into string []
		}
	}
	result1 := strings.Join(mystr, "") //convert the []string as proper string
	result2 := Reverse(result1)        //Call the reverse function and store the return value in variable
	fmt.Printf(result2)                //print the string value of the variable
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
