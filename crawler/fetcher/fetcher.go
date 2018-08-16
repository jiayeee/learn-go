package fetcher

import (
	"bufio"
	"fmt"
	"github.com/gopm/modules/log"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * 猜测 html 的编码类型
 */
func determineEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, e := r.Peek(1024)
	if e != nil {
		log.Error("fetcher err, %v", e)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : ", resp.StatusCode)
	}

	// gbk 转中文乱码解决
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(reader)
}
