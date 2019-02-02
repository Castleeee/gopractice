/* author:owl date:2019/1/27 description: */
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"runtime"
	"strconv"
)

func testif() {

	if contents, err := ioutil.ReadFile("a.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(contents))
	}
	//另一种方法
	var content, errors = ioutil.ReadFile("a.txt")
	if errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println(content)
	}

}

func testswitch(a, b int, op string) int {
	fmt.Println(a, b, op)
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
		fallthrough
	default:
		fmt.Println(a, b, "unknow operator")
	}
	//也可以这样
	switch {
	case op == "+":
		result = a + b
	case op == "-":
		result = a - b
		fallthrough
	default:
		fmt.Println(a, b, "unknow operator")
	}
	return result
}
func testfor(n int) string {

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func functest(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
func functest2(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	testif()
	testswitch(1, 4, "-")
	fmt.Println(testfor(50))
	a, b := functest(1, 2)
	fmt.Println(a, b)
	fmt.Println(apply(pow, a, b))
	fmt.Println(apply(func(a int, b int) int { return int(math.Pow(float64(a), float64(b))) }, a, b))
}
