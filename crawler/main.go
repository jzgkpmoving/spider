package main

import (
	parsera "spider/crawler/aika/parsera"
	"spider/crawler/engine"
	"spider/crawler/persiser"
	"spider/crawler/scheduler"
)

func main() {
	item, err := persiser.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100, ItemSaver: item, RequestProcessor: engine.Worker}
	e.Run(engine.Request{Url: "https://www.xcar.com.cn/", Parser: engine.NewFuncParser(parsera.ParseCarType, "ParseCarType")})
}
