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

	g := make([][]int, n+1)
	result := []int{}

	for i := 0; i < n-1; i++ {
		a, b := NextI2()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 1; i <= n; i++ {
		SortI(g[i])
	}

	var dfs func(int, int)
	dfs = func(u, v int) {
		result = append(result, v)
		for _, w := range g[v] {
			if w != u {
				dfs(v, w)
				result = append(result, v)
			}
		}
	}

	dfs(0, 1)

	OutSlice(result, " ")
}

func main() {
	defer Flush()
	solve()
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

func NextI3() (int, int, int) {
	return NextI(), NextI(), NextI()
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

func OutSlice(slice []int, sep string) {
	for _, v := range slice {
		_, e := fmt.Fprint(writer, v, sep)
		if e != nil {
			panic(e)
		}
	}
	_, e := fmt.Fprintln(writer)
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
