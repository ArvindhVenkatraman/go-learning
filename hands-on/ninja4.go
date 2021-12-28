package main

import "fmt"

func main() {
	//x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(x[:5]) // : starting point will let you get everything upto position 5
	//fmt.Println(x[5:]) // : ending point will let you get everything from that position number till the end of the slice
	//
	//for _, v := range x {
	//	fmt.Println(v)
	//}
	//fmt.Printf("%T\n", x) //print type of the variable
	//
	//e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//e = append(e, 13)
	//fmt.Println(e)
	//e = append(e, 14, 15, 16)
	//fmt.Println(e)
	//f := []int{100, 200, 300}
	//f = append(e, f...) //append with three dots means take all of values from f and append it to e
	//fmt.Println(f)
	//
	//var arsum int32
	//ar := []int32{1, 2, 3, 4, 10, 11}
	//for i := 0; i < len(ar); i++ {
	//	arsum = arsum + ar[i]
	//
	//}
	//fmt.Println(arsum)

	//var pp1 int32
	//var pp2 int32
	//p1 := []int32{17, 28, 30}
	//p2 := []int32{99, 16, 8}
	//fmt.Println("initial pp2", pp2)
	//for i := 0; i < len(p1); i++ {
	//	if p1[i] < p2[i] {
	//		pp2++
	//	} else {
	//		pp1++
	//	}
	//
	//}
	//fmt.Println(pp1, pp2)
	//var arsum1 int64
	//
	//ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	//
	//for i := 0; i < len(ar); i++ {
	//	arsum1 = arsum1 + ar[i]
	//}
	//fmt.Println(arsum1)

	m1 := []int32{11, 2, 4}
	m2 := []int32{4, 5, 6}
	m3 := []int32{10, 8, -12}

	var ld int32
	var rd int32
	for i := 0; i < len(m1); i++ {
		ld = m1[0] + m2[1] + m3[2]
		rd = m1[2] + m2[1] + m3[0]
	}
	fmt.Println(ld)
	fmt.Println(rd)
	tld := ld - rd
	fmt.Println(tld)
}
