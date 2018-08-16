package real

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, e := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if e != nil {
		panic(e)
	}

	s := string(result)
	ioutil.WriteFile("aa.html", result, os.ModeDevice)
	return s
}
