package main

import (
	"flag"
	"log"
	"net/rpc"
	parsera "spider/crawler/aika/parsera"
	"spider/crawler/engine"
	"spider/crawler/scheduler"
	"spider/crawler_distribute/config"
	itemsaver "spider/crawler_distribute/persiser/client"
	"spider/crawler_distribute/rpcsupport"
	worker "spider/crawler_distribute/worker/client"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	item, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100, ItemSaver: item, RequestProcessor: processor}
	e.Run(engine.Request{"https://www.xcar.com.cn/", engine.NewFuncParser(parsera.ParseCarType, config.ParseCarType)})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("error connecting to %s : %v", h, err)
		} else {
			clients = append(clients, client)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
