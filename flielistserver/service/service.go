/* author:owl date:2019-02-03 description:nil */
package service

import (
	"io/ioutil"
	"net/http"
	"os"
)

func Ser(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	con, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(con)
	return nil
}
