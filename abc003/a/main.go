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

func solve() {
	n := NextI()

	Outln(n*5000 + 5000)
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
