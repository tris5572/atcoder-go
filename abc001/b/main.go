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
	m := NextI()

	if m < 100 {
		Out("00")
	} else if m <= 5000 {
		Out(fmt.Sprintf("%02d", m/100))
	} else if m <= 30000 {
		Out(m/1000 + 50)

	} else if m <= 70000 {
		Out(((m/1000)-30)/5 + 80)
	} else {
		Out("89")
	}
}

func main() {
	defer func() {
		e := writer.Flush()
		if e != nil {
			panic(e)
		}
	}()
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

func Out(v ...interface{}) {
	_, e := fmt.Fprintln(writer, v...)
	if e != nil {
		panic(e)
	}
}
