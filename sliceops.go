/* author:owl date:2019/1/27 description:nil */
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func a() {
	var a []int
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 10, 32)
	copy(s2, s1)
	fmt.Println(s2, a)
}
func s(a []int) {
	a[0] = 0
	return

}
func maptest() {
	m := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}
	m1 := make(map[string]int)
	var m2 map[int]string
	fmt.Println(m, m1, m2)
	for k, v := range m {
		fmt.Println(k, v)
	}
	if name, exist := m["k"]; exist {
		fmt.Println(name)
	} else {
		fmt.Println(exist)
	}
}
func Longest_nonrepeating_string(str string) (s string, l int) {
	record := map[string]int{"": 0}
	for i := range s {
		for k, v := range record {
			if strings.Index(k, string(i)) != -1 {
				v += 1
				k = k + string(i)
			} else {
				record["i"] = 1
			}
		}
	}
	return
}
func main() {
	a()
	sn := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(sn[1:3]))
	s(sn)
	fmt.Println(sn)
	fmt.Printf("%T", sn)
	maptest()
	fmt.Println(Longest_nonrepeating_string("abcbbasbasajj"))
}
