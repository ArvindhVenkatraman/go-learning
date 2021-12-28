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
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */
var ld int32

var rd int32
var result1 int32

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	//for i := 0; i < len(arr); i++ {
	//	ld = ld + arr[i][i]
	//}

	// initilize
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		ld = ld + arr[i][i]
		rd = rd + arr[i][j]
	}
	//i := len(arr) - 1
	//fmt.Println("len minus 1", i, len(arr))
	//for range arr {
	//	fmt.Println("ld", arr[i][i])
	//	//rd = rd + arr[i][i]
	//	i--
	//}
	return ld - rd
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	result1 = diagonalDifference(arr)
	fmt.Println(result1)

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
