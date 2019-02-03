/* author:owl date:2019-02-03 description:nil */
package main

func dem() {
	c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
}

func main() {
	dem()
}
