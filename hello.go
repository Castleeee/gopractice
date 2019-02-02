package main

import (
	"fmt"
	"math"
)

func varible() {
	var a int
	var b string
	d, e, f, g := 1, "ss", true, 5
	fmt.Println(a, b, d, e, f, g)
	//var c int32
	//fmt.Println(c+g)
}

func trangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func consta() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}
func ConstsandEnum() {
	const (
		a = iota
		b
		c
		_
		d
	)
	fmt.Println(a, b, c, d)
}
func main() {
	fmt.Println("helloworld")
	varible()
	trangle()
	consta()
	ConstsandEnum()
}
