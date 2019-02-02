/* author:owl date:2019-02-02 description:nil */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Go(wg *sync.WaitGroup, index int) {
	a := 1

	for i := 0; i < 1000000000; i++ {
		a += 1
	}
	fmt.Println(a)
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	a, b := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case v1, ok := <-a:
				if !ok {
					break
				}
				fmt.Println(v1)
			case v2, ok := <-b:
				if !ok {
					break
				}
				fmt.Println(v2)
			}
		}
	}()

	a <- 1
	b <- "wdsdd"
	a <- 3
	b <- "'ss"

	close(a)
	close(b)
	//wg:=sync.WaitGroup{}
	//wg.Add(10)
	//for i:=0;i<10;i++{
	//    go Go(&wg,i)
	//}
	//wg.Wait()
}

//func main() {
//	c := make(chan bool)
//	go func() {
//		fmt.Println("gogogo")
//		c <- true
//	}()
//	a := <-c
//	fmt.Println(a)
//}
