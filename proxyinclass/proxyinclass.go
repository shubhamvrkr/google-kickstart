package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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


func readIntArray(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = readInt()
	}
	return a
}

func main() {

	N := readInt()
	A:=readIntArray(N)
  sort.Ints(A)
	A=unique(A)
	missingroll:=missingstudents(A,N)
	fmt.Printf("%d",missingroll[0])
	for i:=1;i<len(missingroll);i++{
			fmt.Printf(" %d ",missingroll[i])
	}
}

func missingstudents(students []int, N int) []int{
		var missing []int
		studentExpected:=1
		for i:=0;i<len(students);{
			if students[i]==studentExpected{
					i++
					studentExpected++
			}else{
					missing = append(missing,studentExpected);
					studentExpected++
			}
		}
		for i:=students[len(students)-1]+1;i<=N;i++{
			missing = append(missing,i)
		}
		return missing
}

func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}
