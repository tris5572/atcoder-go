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
	// solve
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
// Input
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

func NextI3() (int, int, int) {
	return NextI(), NextI(), NextI()
}

func NextI4() (int, int, int, int) {
	return NextI(), NextI(), NextI(), NextI()
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

func NextSN(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = NextS()
	}
	return a
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Output
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// 出力して改行しない。
func Out(v ...interface{}) {
	_, e := fmt.Fprint(writer, v...)
	if e != nil {
		panic(e)
	}
}

// 出力して改行する。
func Outln(v ...interface{}) {
	_, e := fmt.Fprintln(writer, v...)
	if e != nil {
		panic(e)
	}
}

// slice を sep 区切りで出力する。
func OutSlice(slice []int, sep string) {
	for _, v := range slice {
		_, e := fmt.Fprintf(writer, "%v%s", v, sep)
		if e != nil {
			panic(e)
		}
	}
	_, e := fmt.Fprintln(writer)
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

// n^m (nのm乗)
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

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Slice
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
func SortI(list []int) {
	sort.Ints(list)
}

func SortIRev(list []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
}

/*
-    -     -     -     -     -     -     -     -     -     -     -     -
スライスの重複削除 int版
-    -     -     -     -     -     -     -     -     -     -     -     -
・登場順を守って返す。
*/
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

// -    -     -     -     -     -     -     -     -     -     -     -     -
// スライスが要素を含むかどうか int版
func ContainsI(s []int, v int) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}

/*
-    -     -     -     -     -     -     -     -     -     -     -     -
優先度付きキュー (Priority Queue) int版
-    -     -     -     -     -     -     -     -     -     -     -     -
準備

	・"container/heap" を import する

使い方

	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

補足

	・デフォルトでは Pop() で最小値が取得される。これを最大値にするには
	　Less() の中の符号を逆にする。

参考

	https://pkg.go.dev/container/heap
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Map
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/*
-    -     -     -     -     -     -     -     -     -     -     -     -
マップからキー(int)の一覧を取得する
-    -     -     -     -     -     -     -     -     -     -     -     -
引数の型を必要に応じて書き換える。
*/
func KeysI(m map[int]int) []int {
	ret := make([]int, 0, len(m))
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// Data Structure
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

/*
-    -     -     -     -     -     -     -     -     -     -     -     -
両端キュー  int版
-    -     -     -     -     -     -     -     -     -     -     -     -
双方向リスト(container/list)をラップして、値を int のみに制限して使いやすくしたもの。
*/

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
