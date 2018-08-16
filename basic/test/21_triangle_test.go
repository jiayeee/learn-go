package main

import (
	"math"
	"testing"
)

func CalcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

/**
 * 表格驱动测试
 */
func TestTrangle(t *testing.T) {

	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{12, 35, 38},
		{30000, 40000, 50000},
		{30000, 40000, 30000},
	}

	for _, tt := range tests {
		if actual := CalcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("CalcTrangle(%d, %d) got %d, expected %d\n", tt.a, tt.b, actual, tt.c)
		}
	}
}
