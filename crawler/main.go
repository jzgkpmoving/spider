package main

import (
	parsera "spider/crawler/aika/parsera"
	"spider/crawler/engine"
	"spider/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100}
	e.Run(engine.Request{"https://www.xcar.com.cn/", parsera.ParseCarType})
}
