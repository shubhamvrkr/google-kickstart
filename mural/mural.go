package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

//defer writer.Flush()

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
func avg(a, b int) int {
	return int(math.Floor(float64((a + b) / 2)))
}

func main() {

	var sum, max, j int
	T := readInt()
	for it := 1; it <= T; it++ {
		sum = 0
		max = 0
		n := readInt()
		arr := convertStringToIntAray(readString(), n)
		if n%2 == 0 {
			j = n/2 - 1
		} else {
			j = n / 2
		}
		for i := 0; i <= j; i++ {
			sum += arr[i]
		}
		max = sum
		for index := 1; index < (n - j); index++ {
			sum = int(math.Abs(float64(sum - arr[index-1] + arr[index+j])))
			if sum > max {
				max = sum
			}
		}
		fmt.Printf("Case #%d: %d", it, max)
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
