package main

import (
	"fmt"
)

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var ln, cnt, temp int
	for _, v := range candles {
		if int(v) >= temp {
			temp = int(v)
			ln = temp
		}
	}

	for _, v := range candles {
		//fmt.Println(temp, v)
		if ln <= int(v) {
			temp = int(v)
			//fmt.Println(temp)
			cnt++
		}
	}
	return int32(cnt)

}

func main() {

	candles := []int32{20, 5, 9, 10, 36, 36}

	result := birthdayCakeCandles(candles)
	fmt.Println(result)

}
