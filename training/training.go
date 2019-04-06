package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "sort"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

//defer writer.Flush()
var MOD uint64 = 1000000007

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
	temp := 0
  minHours:= 0;
	lastComputedHours := 0
	T := readInt()
	for it := 1; it <= T; it++ {
		N:=readInt()
		P:=readInt()
		S:=readIntArray(N)
		//O(N log N)
	  sort.Ints(S)
		minHours = getHoursToMakeEqual(S[0:P])
		lastComputedHours = minHours
		// O(N)
		for i:=1;i<N-P+1;i++{
				temp = lastComputedHours -  (S[i+P-2]-S[i-1]) + (P-1)*(S[i+P-1]-S[i+P-2])
				if temp < minHours {
						minHours=temp
				}
				lastComputedHours = temp
		}
		fmt.Printf("Case #%d: %d", it,minHours)
		fmt.Println()
	}
}

func getHoursToMakeEqual(S []int) int{
	max :=S[len(S)-1]
	hours:=0
	for i:=0;i<len(S);i++{
		hours+=max-S[i]
	}
	return hours
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

func quickpow(x uint64, n uint64) uint64 {

	if n == 0 {
		return uint64(1)
	}
	if n == 1 {
		return x
	}
	v := quickpow(x, n/2)
	if n%2 == 1 {
		return v * v % MOD * x % MOD
	} else {
		return v * v % MOD * 1 % MOD
	}
}
