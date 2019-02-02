/* author:owl date:2019/1/27 description:nil */
package main

import "fmt"

func pointertester() {
	a := 3
	pointa := &a
	fmt.Println(*pointa)
}

func slicetest() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a[2:6], a[2:], a[:6], a[:])
}
func changeslice(a [9]int) {
	a[2] = 0
}
func main() {
	pointertester()
	var array1 [5]int
	arrar2 := [3]int{1, 2, 3}
	array3 := [...]int{1, 2, 3, 5, 4}
	var grid [4][5]bool
	fmt.Println(array1, arrar2, array3, grid)
	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}
	for i := range array3 {
		fmt.Println(i)
	}
	for i, v := range array3 {
		fmt.Println(i, v)
	}
	for _, v := range array3 {
		fmt.Println(v)
	}
	slicetest()
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(a)
	s1 := a[2:6]
	s2 := a[3:5]
	fmt.Println(s1, s2)
	fmt.Println(s2[4])

}
