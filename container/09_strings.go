package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	s := "Yes我爱慕课网!"
	fmt.Println(len(s))

	for _, ch := range []byte(s) {

		// 打印十六进制 ASIC 码
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("[%d, %X]", i, ch)
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		fmt.Printf("%X %d %X %c\n", bytes, size, ch, ch)
		bytes = bytes[size:]
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("[%d %c] ", i, ch)
	}

	fmt.Println()
	fmt.Println("--------------------------------")
	fmt.Println(strings.Fields("abc def g h"))
	fmt.Println(strings.Split("abc,def,g,h", ","))
	fmt.Println(strings.Join([]string{"ab", "Cd", "e", "f"}, ","))

	fmt.Println(strings.Contains("abcdef", "de"))
	fmt.Println(strings.Index("abcdef", "de"))

	fmt.Println(strings.ToUpper("abdDe"))
	fmt.Println(strings.ToLower("abdDe"))

	fmt.Println(strings.Trim("aab abdD abe ab a", "abd"))
	fmt.Println(strings.TrimLeft("aab abdD abe ab a", "abd"))
	fmt.Println(strings.TrimRight("aab abdD abe ab a", "abd"))
}
