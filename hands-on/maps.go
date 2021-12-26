package main

import "fmt"

func main() {
	mapvalue := map[string]int{
		"james": 32,
		"bond":  33,
		"M":     21,
	}
	if v, ok := mapvalue["M"]; ok { //comma okay ediom
		fmt.Println("Print Something", v)
	}
	fmt.Println(mapvalue["james"])
}
