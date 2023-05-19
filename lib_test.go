package main

import (
	"fmt"
	"testing"
)

func TestUniqueSliceI(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{name: "何もしない", data: []int{0, 1, 3, 2}, want: []int{0, 1, 3, 2}},
		{name: "通常1", data: []int{0, 1, 1, 2, 3, 3, 4, 1, 2}, want: []int{0, 1, 2, 3, 4}},
		{name: "通常2", data: []int{9, 1, 9, 1, 9}, want: []int{9, 1}},
		{name: "通常3", data: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	isSameSlice := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueSliceI(tt.data)
			if !isSameSlice(got, tt.want) {
				t.Errorf("got= %v, want= %v", got, tt.want)
			}
		})
	}
}

func TestDequeI(t *testing.T) {
	deq := NewDequeI()
	for i := 0; i < 10; i++ {
		deq.PushFront(i)
	}
	fmt.Println(deq)
}
