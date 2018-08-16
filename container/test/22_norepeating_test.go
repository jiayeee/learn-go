package test

import "testing"

func InternationalMaxLengthOfNoneRepeatingSubStr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccured[ch] = i
	}

	return maxLength
}

/**
 * go test -bench . -cpuprofile cpu.out
 * go tool pprof cpu.out
 *		web
 */
func TestNonRepeating(t *testing.T) {

	tests := []struct {
		s   string
		ans int
	}{
		// normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcabcabcd", 4},

		// chinese support
		{"Yes我爱慕课网！", 9},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		if len := InternationalMaxLengthOfNoneRepeatingSubStr(tt.s); len != tt.ans {
			t.Errorf("InternationalMaxLengthOfNoneRepeatingSubStr(%s) got %d, expected %d\n", tt.s, len, tt.ans)
		}
	}
}

func BenchmarkNonRepeating(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		len := InternationalMaxLengthOfNoneRepeatingSubStr(s)
		if len != ans {
			b.Errorf("InternationalMaxLengthOfNoneRepeatingSubStr(%s) got %d, expected %d\n", s, len, ans)
		}
	}
}
