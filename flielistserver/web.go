/* author:owl date:2019-02-03 description:nil */
package main

import (
	"flielistserver/service"
	"fmt"
	"net/http"
)

type serhandler func(writer http.ResponseWriter, request *http.Request) error

func errwapper(ser serhandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := serhandler(writer, request)
		if err != nil {

		}
	}
}
func main() {
	http.HandleFunc("/list/", service.Ser)
	err := http.ListenAndServe(":8888", nil)
	fmt.Println(err)
}
