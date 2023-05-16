package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer Flush()

	n := NextI()
	a := NextIN(n)

	var max1i, max1v, max2i, max2v int

	for i, v := range a {
		if max1v < v {
			max2i = max1i
			max2v = max1v
			max1i = i + 1
			max1v = v
		} else if max2v < v {
			max2i = i + 1
			max2v = v
		}
	}

	Out(max2i)

}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// I/O
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func init() {
	scanner.Buffer([]byte{}, math.MaxInt64)
	scanner.Split(bufio.ScanWords)
}

func NextI() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func NextI2() (int, int) {
	return NextI(), NextI()
}

func NextIN(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = NextI()
	}
	return a
}

func NextF() float64 {
	scanner.Scan()
	f, e := strconv.ParseFloat(scanner.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func NextS() string {
	scanner.Scan()
	return scanner.Text()
}

func Out(v ...interface{}) {
	_, e := fmt.Fprintln(writer, v...)
	if e != nil {
		panic(e)
	}
}

func Flush() {
	e := writer.Flush()
	if e != nil {
		panic(e)
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Number
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func MaxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinI(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func AbsI(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Slice
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func SortI(list []int) {
	sort.Ints(list)
}

func SortIR(list []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
}
