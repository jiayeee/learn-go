package engine

import (
	"fmt"
	"github.com/gopm/modules/log"
	"learn-go/crawler/fetcher"
)

type SimpleEngine struct {
}

func worker(r Request) (ParserResult, error) {
	fmt.Printf("fetching url : %s\n", r.Url)

	contents, e := fetcher.Fetch(r.Url)
	if e != nil {
		log.Error("fetcher error, url : %s", r.Url)

		return ParserResult{}, e
	}

	return r.ParserFunc(contents), nil
}

func (se SimpleEngine) Run(seeds ...Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	var count int
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, e := worker(r)
		if e != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		fmt.Println(len(parserResult.Items))
		for _, item := range parserResult.Items {
			count++
			fmt.Printf("got item : %v, count = #%d#\n", item, count)
		}
	}
}
