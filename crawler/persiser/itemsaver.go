package persiser

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
	"spider/crawler/engine"
)

func ItemSaver() (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	go func(client *elastic.Client) {

		count := 1
		for {
			item := <-out
			log.Printf("got %d Item %v", count, item)
			count++
			err := save(item, client)
			if err != nil {
				log.Printf("Item %d Saver error %v", count, err)
			}
		}
	}(client)
	return out, err
}

func save(item engine.Item, client *elastic.Client) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return err

}
