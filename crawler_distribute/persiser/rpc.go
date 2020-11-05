package persiser

import (
	"github.com/olivere/elastic/v7"
	"log"
	"spider/crawler/engine"
	"spider/crawler/persiser"
)

type ItemSaverService struct {
	Client *elastic.Client
}

func (s *ItemSaverService) Save(item engine.Item, reslut *string) error {
	err := persiser.Save(item, s.Client)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*reslut = "ok"
	} else {
		log.Printf("Error saving item %v %v", item, err)
	}
	return err
}
