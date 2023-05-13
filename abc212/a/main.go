package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	a, b := nexti2()

	if 0 < a && b == 0 {
		out("Gold")
	} else if a == 0 && 0 < b {
		out("Silver")
	} else {
		out("Alloy")
	}
}

func init() {
	scanner.Buffer([]byte{}, math.MaxInt64)
	scanner.Split(bufio.ScanWords)
	if len(os.Args) > 1 {
		b, e := ioutil.ReadFile("./input1")
		if e != nil {
			panic(e)
		}
		scanner = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}

func nexti() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nexti2() (int, int) {
	return nexti(), nexti()
}

func out(v ...interface{}) {
	_, e := fmt.Fprintln(writer, v...)
	if e != nil {
		panic(e)
	}
}

func flush() {
	e := writer.Flush()
	if e != nil {
		panic(e)
	}
}
