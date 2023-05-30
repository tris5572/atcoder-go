package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	a := NextS()
	b := NextS()

	result := true

	for i := range a {
		if a[i] == b[i] {
			continue
		}
		if a[i] == '@' && strings.ContainsRune("atcoder", rune(b[i])) {
			continue
		}
		if b[i] == '@' && strings.ContainsRune("atcoder", rune(a[i])) {
			continue
		}
		result = false
		break
	}

	if result {
		Outln("You can win")
	} else {
		Outln("You will lose")
	}
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

func NextS() string {
	scanner.Scan()
	return scanner.Text()
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
