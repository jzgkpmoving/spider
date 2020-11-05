package client

import (
	"log"
	"spider/crawler/engine"
	"spider/crawler_distribute/config"
	"spider/crawler_distribute/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {

		count := 1
		for {
			item := <-out
			log.Printf("got %d Item %v", count, item)
			count++

			result := ""
			err := client.Call(config.ItemSaverServer, item, &result)
			if err != nil {
				log.Printf("Item %d Saver error %v", count, err)
			}
		}
	}()
	return out, err
}
