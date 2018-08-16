package citylist

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {

	//contents, e := fetcher.Fetch("http://album.zhenai.com/u/108297578")
	//fmt.Println(contents, e)
	contents, _ := ioutil.ReadFile("profile_test_data.html")

	parserResult := ParseProfile(contents, "李四")
	fmt.Println(parserResult.Items)
}
