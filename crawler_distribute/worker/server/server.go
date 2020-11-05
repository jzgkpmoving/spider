package main

import (
	"flag"
	"fmt"
	"log"
	"spider/crawler_distribute/config"
	"spider/crawler_distribute/rpcsupport"
	"spider/crawler_distribute/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServRpc(config.WorkerPort0, worker.CrawlService{}))
}
