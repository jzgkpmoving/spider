package worker

import (
	"spider/crawler/engine"
)

type CrawlService struct {
}

func (CrawlService) Process(request Request, result *ParseResult) error {
	enginereq, err := DeserializeRequest(request)
	if err != nil {
		return nil
	}
	engineReslut, err := engine.Worker(enginereq)

	if err != nil {
		return err
	}
	*result = SerialResult(engineReslut)
	return nil
}
