package citylist

import (
	"learn-go/crawler/engine"
	"regexp"
)

// ^> 表示非大于号
var cityListReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-b][a-z0-9]+)"[^>]+>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParserResult {
	matches := cityListReg.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
