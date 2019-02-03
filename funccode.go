package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()       //取下一个元素
	if next > 10000 { //达到10000以上结束
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next) //转换成字符串

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

//把里面的内容打印出来
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = Fibonacci()
	printFileContents(f)
}

//type的作用https://blog.csdn.net/hzwy23/article/details/79890778
