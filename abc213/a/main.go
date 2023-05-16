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

func main() {
	defer Flush()

	a, b := NextI2()

	Out(a ^ b)
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
// Numbers
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
