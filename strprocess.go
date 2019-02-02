/* author:owl date:2019/1/27 description:nil */
package main

import (
	"fmt"
)

type A struct {
	s int
	b float32
}

func (a A) increase() int {
	a.s += 100
	return 1
}
func main() {
	s := "yes我坎姆科湾" //utf-8编码
	for v, b := range []byte(s) {
		fmt.Printf("%X", b)
		fmt.Println("|", v)

	}
	a := A{100, 8.766}
	var p1 = &a
	var p2 *A = &a
	fmt.Println(p1, p2)
	fmt.Println()
	a.increase()
	fmt.Println(a.s, a.increase(), a.s)
}
