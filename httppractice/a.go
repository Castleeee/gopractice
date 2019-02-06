/* author:owl date:2019-02-03 description:nil */
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	url := "https://www.imooc.com/"

	//res,err:=http.Get(url)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Mobile Safari/537.36")
	request.Header.Add("Pragma", "no-cache")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(via, req)
			return nil
		},
	}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", body)
}
