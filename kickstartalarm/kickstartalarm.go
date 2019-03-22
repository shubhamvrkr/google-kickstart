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

func avg(a, b int) int {
	return int(math.Floor(float64((a + b) / 2)))
}

func main() {
	var A []int
	T := readInt()
	for it := 1; it <= T; it++ {

		N := readInt()
		K := readInt()
		powersummation := make([]uint64, K)
		x1 := readInt()
		y1 := readInt()
		C := readInt()
		D := readInt()
		E1 := readInt()
		E2 := readInt()
		F := readInt()

		A = computeParameterArray(N, K, x1, y1, C, D, E1, E2, F)

		for k := 0; k < K; k++ {
			powersummation[k] = computeKthPower(A, N, k+1)
		}
		sum := uint64(0)
		for i := 0; i < K; i++ {
			sum += powersummation[i]
		}
		sum = sum % 1000000007
		fmt.Printf("Case #%d: %d", it, sum)
		fmt.Println()
	}
}
func computeKthPower(A []int, N int, K int) uint64 {
	memo := create2DArrayUnsigned(N, N)
	powersum := uint64(0)
	for i := 0; i < N; i++ {
		memo[i][i] = uint64(A[i])
		powersum += uint64(A[i])
	}
	for i := 1; i < N; i++ {
		for j := 0; j < N-i; j++ {
			z := j + i
			value := memo[j][z-1] + uint64(A[z])*quickpow(uint64(i+1), uint64(K))
			memo[j][z] = value
			powersum += value
		}
	}
	return powersum % 1000000007
}

func computeParameterArray(N int, K int, x1 int, y1 int, C int, D int, E1 int, E2 int, F int) []int {
	A := make([]int, N)
	X := make([]int, N)
	Y := make([]int, N)
	X[0] = x1
	Y[0] = y1
	A[0] = (X[0] + Y[0]) % F
	for i := 1; i < N; i++ {
		X[i] = (C*X[i-1] + D*Y[i-1] + E1) % F
		Y[i] = (D*X[i-1] + C*Y[i-1] + E2) % F
		A[i] = (X[i] + Y[i]) % F
	}
	return A
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
