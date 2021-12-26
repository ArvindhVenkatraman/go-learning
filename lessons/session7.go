package main

import (
	"fmt"
	"strings"
)

var mystr []string

func main() {
	str := "34o3ll435e  39h   "
	str = strings.TrimSpace(str)
	for _, v := range str {
		if v > 65 {
			//fmt.Println(v, string(v))
			mystr = append(mystr, string(v))
		}
	}
	result1 := strings.Join(mystr, "")
	result2 := Reverse(result1)
	fmt.Printf(string(result2))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
