package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// defer 列表为先进后出，所以这里会先打印出2，后打印出1
// 使用 defer 的场景：Open/Close、Lock/Unlock、PrintHeader/PrintFooter
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	//panic("error occured")

	return

	fmt.Println(4)
}

func writeFile(filename string) {
	file, e := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	e = errors.New("this is a custom error")

	if e != nil {
		if pathError, ok := e.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)

			return
		} else {
			panic(e)
		}
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "abcde")
}

func main() {
	//tryDefer()

	writeFile("defer.txt")
}
