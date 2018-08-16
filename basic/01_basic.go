package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 * 内建变量类型

 * bool string
 *
 * (u)int (u)int8 (u)int16 (u)int32 (u)int64 (u)intptr
 *
 * byte rune
 *
 * float32 float64 complex64 complex128
 *
 *
 *
 */
var aa = 3
var bb = true
var ss = "kkk"

var (
	cc = 4
	dd = false
	ee = "eeeee"
)

func variableZeroValue() {
	var i int
	var s string
	fmt.Printf("%d, %q\n", i, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableInitialDeduction() {
	var a, b, c, s = 1, 3, true, "def"
	fmt.Println(a, b, c, s)
}

func variableInitialShorter() {

	// := 变量定义只能在包内部
	a, b, c, s := 1, 3, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

/**
 * 没有隐式转换，只能强制类型转换
 */
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 1, 3
	c := math.Sqrt(a*a + b*b)

	const (
		f    = "c.txt"
		d, e = 3, 6
	)

	fmt.Printf("%s, %f\n", filename, c)
	fmt.Println(f, d, e)
}

func enums() {
	const (
		java   = 0
		python = 1
		golang = 2
	)

	const (
		A = iota
		B
		C
		_
		D
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(java, python, golang)
	fmt.Println(A, B, C, D)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")

	variableZeroValue()
	variableInitialValue()
	variableInitialDeduction()
	variableInitialShorter()
	euler()
	triangle()
	consts()
	enums()
}
