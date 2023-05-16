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

// 座標圧縮で解く。
func solve() {
	_, _, n := NextI3()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = NextI2()
	}

	// 出現する値のリスト(ソート済み)を作成する
	numsA := UniqueSliceI(a)
	SortI(numsA)
	numsB := UniqueSliceI(b)
	SortI(numsB)

	// 圧縮したデータを作成する
	compA := map[int]int{}
	compB := map[int]int{}

	index := 1
	for _, k := range numsA {
		compA[k] = index
		index++
	}

	index = 1
	for _, k := range numsB {
		compB[k] = index
		index++
	}

	for i := range a {
		Out(compA[a[i]], compB[b[i]])
	}
}

func main() {
	defer Flush()
	solve()
}

func UniqueSliceI(a []int) []int {
	m := map[int]bool{}
	r := make([]int, 0)
	for _, v := range a {
		if !m[v] {
			m[v] = true
			r = append(r, v)
		}
	}
	return r
}

func KeysI(m map[int]int) []int {
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
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
