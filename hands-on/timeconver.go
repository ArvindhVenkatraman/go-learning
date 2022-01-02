package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here

	s = "07:05:45PM"
	//layout := "19:05:45"
	t, _ := time.Parse("03:04:05PM", s)
	s = t.Format("15:04:05")
	return s

}

func main() {
	var s string
	result := timeConversion(s)
	fmt.Println(result)

}
