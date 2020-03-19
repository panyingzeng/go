// map project main.go
package main

import (
	"fmt"
)

//寻找最长不重复的子串
func lengthOfNoRepeatedSubstr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	mlength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > mlength {
			mlength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return mlength
}

func main() {
	m := map[string]string{ //map是无序的，hashMap
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2==empty map
	var m3 map[string]int      //m3==nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing Map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//拼错情况下
	cName, ok := m["cause"]
	fmt.Println(cName, ok) //会打出一个空串，下接
	//因为go语言会有一个zeroValue
	if cName, ok := m["cause"]; ok { //ko判是否为空
		fmt.Println(cName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Daleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	//调用lengthOfNoRepeatedSubstr()
	fmt.Println(lengthOfNoRepeatedSubstr("abcab"))
	fmt.Println(lengthOfNoRepeatedSubstr(""))
	fmt.Println(lengthOfNoRepeatedSubstr("a"))
	fmt.Println(lengthOfNoRepeatedSubstr("bbbbb"))
	fmt.Println(lengthOfNoRepeatedSubstr("abcdef"))
	fmt.Println(lengthOfNoRepeatedSubstr("一二三四五六七"))
	fmt.Println(lengthOfNoRepeatedSubstr("哈哈哈"))

}
