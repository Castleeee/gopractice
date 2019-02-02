/* author:owl date:2019-02-02 description:nil */
package main

import "fmt"

func fb() func(x int ) {
    a:=make([]int,0)
    a=append(a,1,1)
    fmt.Println(a)
    //a[0],a[1]=1,1
    return func (x int ){
        if x<=2{
            fmt.Println(a[x])
            return
        }
        for i:=2;i<x-1;i++{
            a=append(a,a[i-1]+a[i-2])
        }
        fmt.Println(a)
    }
    
    
}
func fb1()func() uint {
    var a,b uint =1,1
    return func() uint {
         a,b =b,a+b
        return a+b
    }
}
func main() {
    x:=fb()
    x(16)
    f:=fb1()
    a:=[...]int{1,2,3}
    for i:=0;i<100;i++ {
        fmt.Println(f())
    }
    for i,_:=range[len(a)]int{}{
        fmt.Println(i)
    }
}
