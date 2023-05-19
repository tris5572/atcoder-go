package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	h, w := NextI2()
	area := make([]int, h*w) // 0が道、1が壁
	cost := make([]int, h*w) // その位置に行くまでの最小コスト
	deque := NewDequeI()

	// i行j列をインデックスへ変換する
	ij2i := func(i, j int) int {
		return i*w + j
	}

	// インデックスをi行j列へ変換する
	i2ij := func(i int) (int, int) {
		return i / w, i % w
	}

	// 座標が範囲内に収まっているかどうかを返す。
	isin := func(y, x int) bool {
		if 0 <= y && 0 <= x && y < h && x < w {
			return true
		}
		return false
	}

	for i := 0; i < h; i++ {
		s := NextS()
		for j, c := range s {
			if c == '.' {
				area[ij2i(i, j)] = 0
			} else {
				area[ij2i(i, j)] = 1
			}
			cost[ij2i(i, j)] = 999999
		}
	}

	cost[ij2i(0, 0)] = 0
	deque.PushBack(0)

	for {
		v, err := deque.PopFront()
		if err != nil {
			break
		}
		a, b := i2ij(v) // a行b列

		// 上下左右のコスト0の移動をチェック
		for i := -1; i <= 1; i++ {
			y := a + i
			for j := -1; j <= 1; j++ {
				// 斜め移動はできないのでチェック対象外
				if AbsI(i) == 1 && AbsI(j) == 1 {
					continue
				}
				x := b + j
				if isin(y, x) {
					z := ij2i(y, x)
					if area[z] == 0 {
						nc := MinI(cost[v], cost[z])
						if nc != cost[z] {
							cost[z] = nc
							deque.PushFront(z)
						}
					}
				}
			}
		}

		// 壁を破壊するコスト1の移動をチェック
		for i := -2; i <= 2; i++ {
			y := a + i
			for j := -2; j <= 2; j++ {
				// 縦横それぞれ +2 or -2 要するに斜めの壁は破壊できないのでチェック対象外
				if AbsI(i) == 2 && AbsI(j) == 2 {
					continue
				}
				x := b + j
				if isin(y, x) {
					z := ij2i(y, x)
					nc := MinI(cost[v]+1, cost[z])
					if nc != cost[z] {
						cost[z] = nc
						deque.PushBack(z)
						// }
					}
				}
			}
		}
	}

	Out(cost[ij2i(h-1, w-1)])
}

func main() {
	defer Flush()
	solve()
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type DequeI struct {
	queue *list.List
}

func NewDequeI() *DequeI {
	d := new(DequeI)
	d.queue = list.New()
	return d
}

func (d *DequeI) Len() int {
	return d.queue.Len()
}

func (d *DequeI) PushFront(v int) {
	d.queue.PushFront(v)
}

func (d *DequeI) PushBack(v int) {
	d.queue.PushBack(v)
}

func (d *DequeI) PopFront() (int, error) {
	el := d.queue.Front()
	if el == nil {
		return 0, fmt.Errorf("deque is empty")
	}
	d.queue.Remove(el)
	return el.Value.(int), nil
}

func (d *DequeI) PopBack() (int, error) {
	el := d.queue.Back()
	if el == nil {
		return 0, fmt.Errorf("deque is empty")
	}
	d.queue.Remove(el)
	return el.Value.(int), nil
}

func (d *DequeI) String() string {
	var b strings.Builder
	first := true
	el := d.queue.Front()
	for el != nil {
		if !first {
			b.WriteString(", ")
		}
		first = false
		b.WriteString(fmt.Sprint(el.Value))
		el = el.Next()
	}

	return b.String()
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
