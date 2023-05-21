package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	n := NextI()
	lines := NextSN(n)
	array := make([]bool, 288) // 雨が降ったかどうかを5分ごと保持する

	for _, line := range lines {
		a := strings.Split(line, "-")
		s := t2i(a[0], false)
		e := t2i(a[1], true)
		for i := s; i <= e; i++ {
			array[i] = true
		}
	}

	before := false
	for i, a := range array {
		if a != before {
			if a {
				Out(i2t(i))
				Out("-")
			} else {
				Outln(i2t(i))
			}
			before = a
		}
		if i == 287 && a {
			Outln("2400")
		}
	}
}

// 時刻をインデックスに変換する
func t2i(s string, end bool) int {
	if s == "2400" {
		return 12*24 - 1
	}
	h, _ := strconv.Atoi(s[0:2])
	m, _ := strconv.Atoi(s[2:4])
	// 分が5の倍数ちょうどのときは、1個前のインデックスにする(出力時に対処するため)
	if end && m%5 == 0 {
		m /= 5
		m -= 1
	} else {
		m /= 5
	}
	return h*12 + m
}

func i2t(i int) string {
	if 288 <= i {
		return "2400"
	}
	return fmt.Sprintf("%02d%02d", i/12, (i%12)*5)
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

func NextS() string {
	scanner.Scan()
	return scanner.Text()
}

func NextSN(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = NextS()
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
