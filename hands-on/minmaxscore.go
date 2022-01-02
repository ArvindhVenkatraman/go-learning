package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int64) {
	//arr = []int32{140638725, 436257910, 953274816, 953274816, 362748590}

	var sum1, sum2 int64
	for i, _ := range arr {
		sum1 = sum1 + arr[i]
		fmt.Println("sum1:arr[i]", sum1, arr[i])
	}

	for j, _ := range arr {
		sum2 = sum1 - arr[j]
		arr[j] = sum2
		//fmt.Println("sum1:arr[j]", sum2, arr[j])
	}

	var sn, ln, temp int

	for _, v := range arr {
		if int(v) >= temp {
			temp = int(v)
			ln = temp
		}
	}
	for _, v := range arr {
		if int(v) <= temp {
			temp = int(v)
			sn = temp
		}
	}

	fmt.Println(ln, sn)

}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int64(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
