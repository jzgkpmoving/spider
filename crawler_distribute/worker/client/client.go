package client

import (
	"net/rpc"
	"spider/crawler/engine"
	"spider/crawler_distribute/config"
	"spider/crawler_distribute/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)

		var sReslut worker.ParseResult

		c := <-clientChan

		err := c.Call(config.CrawlServiceRpc, sReq, &sReslut)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sReslut), nil
	}
}
