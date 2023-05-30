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

func solve() {
	n := NextI()
	k := NextI()
	r := NextIN(n)

	rate := 0.0

	// ソートして、大きい方からk個のうち、小さい方から見ていく。
	SortI(r)
	for i := n - k; i < n; i++ {
		rate = (float64(r[i]) + rate) / 2
	}

	Outln(rate)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func main() {
	defer func() {
		e := writer.Flush()
		if e != nil {
			panic(e)
		}
	}()
	solve()
}

func init() {
	scanner.Buffer([]byte{}, math.MaxInt64)
	scanner.Split(bufio.ScanWords)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// I/O
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func NextI() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func NextIN(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = NextI()
	}
	return a
}

func Out(v ...interface{}) {
	_, e := fmt.Fprint(writer, v...)
	if e != nil {
		panic(e)
	}
}

func Outln(v ...interface{}) {
	_, e := fmt.Fprintln(writer, v...)
	if e != nil {
		panic(e)
	}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func SortI(list []int) {
	sort.Ints(list)
}
