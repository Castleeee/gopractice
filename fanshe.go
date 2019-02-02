/* author:owl date:2019/2/1 description:nil */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) hello() {
	fmt.Println("hello")
}

type Admin struct {
	User //
	fifa int
}

func main() {
	fmt.Printf("a")
	user := User{2, "xiaoming", 15}
	a := 1
	info(user)
	info(a)
	admin := Admin{user, 123}
	t := reflect.TypeOf(admin)
	fmt.Println(t.Field(0), t.FieldByIndex([]int{0, 1}))
	k := reflect.ValueOf(&a)
	k.Elem().SetInt(999)
	fmt.Println(a)

	change(&user)
	fmt.Println(user)

}

func change(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		return
	} else {
		v = v.Elem() //注意类型相同

	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("bad")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("bybe")
	}

}

func info(o interface{}) {

	t := reflect.TypeOf(o)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("not my fault")
		return
	}
	fmt.Println(t.Name())
	f := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		x := t.Field(i)
		y := f.Field(i).Interface()
		fmt.Println(x.Name, x.Type, y)
	}
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i))
	}
}
