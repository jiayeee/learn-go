package main

import "fmt"

/**
 * map[K]V
 * map[K]map{K]V
 *
 * 哪些类型可以作为 map 的 key?
 * map 使用哈希表，必须可以比较大小
 * 除了 slice/map/function 的内建类型都可以作为 key
 * struct 类型不包含上述字段也可以作为 key
 */
func main() {

	m := map[string]string{
		"name":    "oceanz",
		"course":  "golang",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // empty map
	var m3 map[string]int      // nil

	fmt.Println(m, m2, m3)
	fmt.Println(m3 == nil)

	fmt.Println("----------------------")
	fmt.Println("traverse map")

	// map 遍历不保证顺序
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("----------------------")
	fmt.Println("get values")
	courseName := m["course"]
	fmt.Println(courseName)

	// 不存在的 key 会返回 zero value，可以通过加 ok 来判断
	causeName := m["cause"]
	fmt.Println(causeName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("missing key")
	}

	fmt.Println("----------------------")
	fmt.Println("delete values")

	delete(m, "name")
	name, ok := m["name"]
	fmt.Println(name, ok)
}
