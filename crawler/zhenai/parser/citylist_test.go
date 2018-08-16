package citylist

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {

	contents, _ := ioutil.ReadFile("citylist_test_data.html")

	list := ParseCityList(contents)

	for _, request := range list.Requests {
		fmt.Println(request.Url)
	}

	for _, item := range list.Items {
		fmt.Println(item)
	}
}
