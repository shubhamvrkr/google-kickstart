package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const (
	//MaxScanTokenSize = 64 * 1024 // package default
	MaxScanTokenSize = 8 * 1024 * 1024 // 8M
)

func init() {
	scanner.Buffer(make([]byte, 0, MaxScanTokenSize), MaxScanTokenSize)
	scanner.Split(bufio.ScanWords)
}

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readInt64() int64 {
	v, _ := strconv.ParseInt(readString(), 10, 64)
	return v
}

func readIntArray(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = readInt()
	}
	return a
}

func readInt64Array(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = readInt64()
	}
	return res
}

func readFloat64() float64 {
	v, _ := strconv.ParseFloat(readString(), 64)
	return v
}

func readFloat64Array(n int) []float64 {
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = readFloat64()
	}
	return res
}

func create2DArraySigned(rows int, columns int) [][]int {
	a := make([][]int, rows)
	for i := range a {
		a[i] = make([]int, columns)
	}
	return a
}

func create2DArrayUnsigned(rows int, columns int) [][]uint64 {
	a := make([][]uint64, rows)
	for i := range a {
		a[i] = make([]uint64, columns)
	}
	return a
}

func main() {

	var N,pathLength,i,j int

	var P string
	var firsthalf,secondhalf string

	T := readInt()
	for it := 1; it <= T; it++ {

		firsthalf = ""
		secondhalf= ""
		N = readInt()
		P = readString()
		i=0;
		j=2*N-3
		pathLength = 2*N-2

		for i=0;i<pathLength/2;i++{
			if P[i] !=P[j]{
				firsthalf=firsthalf+string(P[j])
				secondhalf=string(P[i])+secondhalf
			}else if P[i]==P[j] && P[i]=='E'{
				firsthalf=firsthalf+"S"
				secondhalf="S"+secondhalf
			}else{
				firsthalf=firsthalf+"E"
				secondhalf="E"+secondhalf
			}
			j--
	  }
		fmt.Printf("Case #%d: %s", it,firsthalf+secondhalf)
		fmt.Println()
	}
}

func convertStringToIntAray(input string, n int) []int {
	arr := make([]int, n)
	for i, char := range input {
		val, _ := strconv.Atoi(string(char))
		arr[i] = val
	}
	return arr
}

func print2DArray(arr [][]uint64, rows int, colums int) {

	for i := 0; i < rows; i++ {
		for j := 0; j < colums; j++ {
			fmt.Printf("%d\t\t", arr[i][j])
		}
		fmt.Println()
	}
}
