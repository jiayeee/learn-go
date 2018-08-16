package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func grade(score int) string {

	g := ""

	switch {
	case score < 0 || score > 100:
		// Sprintf 将格式化后的字符串赋值给一个变量
		panic(fmt.Sprintf("wrong score : %d\n", score))
	case score < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

/**
 * 整数转二进制
 */
func convertToBin(n int) string {

	if n == 0 {
		return "0"
	}

	result := ""
	for ; n > 0; n = n / 2 {
		lsb := n % 2

		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
 * 省略初始条件和递增条件的循环
 */
func printFile(filename string) {
	file, e := os.Open(filename)
	if e != nil {
		panic(e)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/**
 * 不带任何条件的循环，即 死循环
 */
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {

	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println("-----------------------------------")

	if bytes, e := ioutil.ReadFile(filename); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%s\n", bytes)
	}
	fmt.Println("-----------------------------------")

	fmt.Println(grade(0), grade(33), grade(77), grade(88), grade(99))
	fmt.Println("-----------------------------------")

	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(322323),
		convertToBin(0))
	fmt.Println("-----------------------------------")

	printFile("abc.txt")
	fmt.Println("-----------------------------------")

	forever()
}
