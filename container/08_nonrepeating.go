package main

import "fmt"

/**
 * 寻找最长不含有重复字符的子串
 *
 */
// 国际版
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

func maxLengthOfNoneRepeatingSubStr(s string) int {
	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

		if lastI, ok := lastOccured[ch]; ok {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccured[ch] = i
	}

	return maxLength
}

func main() {

	fmt.Println(maxLengthOfNoneRepeatingSubStr("abc"))
	fmt.Println(maxLengthOfNoneRepeatingSubStr("abcddeed"))
	fmt.Println(maxLengthOfNoneRepeatingSubStr("bbbbbbbbbbbb"))

	fmt.Println(maxLengthOfNoneRepeatingSubStr(""))
	fmt.Println(maxLengthOfNoneRepeatingSubStr("abcdefg"))

	fmt.Println(InternationalMaxLengthOfNoneRepeatingSubStr("Yes我爱慕课网！"))
	fmt.Println(InternationalMaxLengthOfNoneRepeatingSubStr("一二三二一"))
}
