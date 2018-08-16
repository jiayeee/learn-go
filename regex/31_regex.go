package main

import (
	"fmt"
	"regexp"
)

const text = `my email is yunheyingzhu@gmail.com
email is abc@def.org
email2 is        kkk@qq.com
email3 is qqq@mail.tencent.com
`

// .+@\.+
// .+ 一个或多个
// .* 零个或多个
// . 表示任何字符，在匹配点的时候要使用 \.
// 在方括号里面，点不用转义
// 小括号表示提取的部分
const regexEmail = `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`

func main() {

	re := regexp.MustCompile(regexEmail)
	as := re.FindAllString(text, -1)

	fmt.Println(as)

	s := re.FindAllStringSubmatch(text, -1)

	for _, m := range s {
		fmt.Println(m)
	}
}
