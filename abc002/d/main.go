package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	n, m := NextI2()

	data := make([][]bool, n)
	for i := 0; i < n; i++ {
		data[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		x, y := NextI2()
		x -= 1
		y -= 1
		data[x][y] = true
		data[y][x] = true
	}

	// ビットで渡されたパターンが全員知り合いかどうかを調べる。
	check := func(bit int) bool {
		targets := []int{}
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) != 0 {
				targets = append(targets, i)
			}
		}

		if len(targets) <= 1 {
			return false
		}

		for i := 0; i < len(targets); i++ {
			for j := i + 1; j < len(targets); j++ {
				if i != j {
					if !data[targets[i]][targets[j]] {
						return false
					}
				}
			}
		}

		return true
	}

	max := 1
	// ビットパターンで全パターンを調べる
	for i := 0; i < PowI(2, n); i++ {
		f := check(i)
		if f {
			c := bits.OnesCount(uint(i))
			max = MaxI(max, c)
		}
	}

	Outln(max)
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

func NextI2() (int, int) {
	return NextI(), NextI()
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

func MaxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PowI(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
