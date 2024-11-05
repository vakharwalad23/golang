package main

import (
	"fmt"
	"testing"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMin(t *testing.T) {
	ans := Min(10, -4)
	if ans != -4 {
		t.Errorf("Min(10, -4) = %d; want -4", ans)
	}
}

func TestMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{10, 2, 2},
		{1, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Min(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Min(%d, %d) = %d; want %d", tt.a, tt.b, ans, tt.want)
			}
		})
	}
}

func BenchmarkMin(b *testing.B) {



	for i := 0; i < b.N; i++ {
		Min(1, 2)
	}
}