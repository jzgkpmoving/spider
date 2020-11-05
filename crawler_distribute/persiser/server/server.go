package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"spider/crawler_distribute/config"
	"spider/crawler_distribute/persiser"
	"spider/crawler_distribute/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(config.ItemSaverHost, config.ElastIndex))
}
func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return err
	}
	return rpcsupport.ServRpc(host, &persiser.ItemSaverService{
		Client: client,
	})
}
