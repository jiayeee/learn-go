package main

import (
	"learn-go/crawler/engine"
	"learn-go/crawler/persist"
	"learn-go/crawler/scheduler"
	"learn-go/crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChannel: persist.ItemSaver(),
	}
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: citylist.ParseCityList,
	})
}
