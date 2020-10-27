package main

import (
	parsera "spider/crawler/aika/parsera"
	"spider/crawler/engine"
)

func main() {
	engine.Run(engine.Request{"https://www.xcar.com.cn/", parsera.ParseCarType})
}
