package main1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	var judgeresponse string
	var guess int
	defer writer.Flush()
	T := readInt()
	for it := 1; it <= T; it++ {
		a, b := readInt(), readInt()
		a = a + 1
		n := readInt()
		for trials := 1; trials <= n; trials++ {
			guess = avg(a, b)
			fmt.Println(guess)
			defer writer.Flush()
			judgeresponse = readString()
			if strings.EqualFold(judgeresponse, "CORRECT") {
				break
			} else if strings.EqualFold(judgeresponse, "TOO_SMALL") {
				a = guess + 1
			} else if strings.EqualFold(judgeresponse, "TOO_BIG") {
				b = guess - 1
			}
		}
	}
}
