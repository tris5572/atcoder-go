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
	defer Flush()

	s := NextS()
	array := strings.Split(s, "")
	flag := true

	n, _ := strconv.Atoi(s)
	if n%1111 == 0 {
		Out("Weak")
		return
	}

	for i := 0; i < 3; i++ {
		a, _ := strconv.Atoi(array[i])
		b, _ := strconv.Atoi(array[i+1])
		if !isWeek(a, b) {
			flag = false
			break
		}
	}
	if flag {
		Out("Weak")
	} else {
		Out("Strong")
	}
}

func isWeek(a, b int) bool {
	if a == 9 && b == 0 {
		return true
	}
	if b-a == 1 {
		return true
	}
	return false
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// I/O
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

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
