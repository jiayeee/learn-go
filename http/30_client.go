package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, e := http.NewRequest(http.MethodGet, "https://www.imooc.com/", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Mobile Safari/537.36")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			fmt.Println("redirect : ", req)
			fmt.Println("----------------------------------------------------------------")

			// 返回 nil 表示 允许重定向，否则不允许重定向
			return nil
		},
	}

	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	r, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(e)
	}

	//fmt.Printf("%s\n", r)
	r = r
}
