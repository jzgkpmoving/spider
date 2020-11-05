package worker

import (
	"errors"
	"log"
	"spider/crawler/aika/parsera"
	"spider/crawler/engine"
	"spider/crawler_distribute/config"
)

type SerializedParser struct {
	FunctionName string
	Args         interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := DeserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{r.Url, parser}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		req, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("DeserializeResult error %v", err)
			continue
		}
		result.Requests = append(result.Requests, req)
	}
	return result
}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FunctionName {
	case config.ParseBrand:
		return engine.NewFuncParser(parsera.ParseBrand, config.ParseBrand), nil
	case config.ParseCar:
		return engine.NewFuncParser(parsera.ParseCarInfo, config.ParseCar), nil
	case config.ParseCarList:
		return engine.NewFuncParser(parsera.ParseCar, config.ParseCarList), nil
	case config.ParseCarType:
		return engine.NewFuncParser(parsera.ParseCarType, config.ParseCarType), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{r.Url, SerializedParser{name, args}}
}

func SerialResult(result engine.ParseResult) ParseResult {
	r := ParseResult{
		Items: result.Items,
	}

	for _, req := range result.Requests {
		r.Requests = append(r.Requests, SerializeRequest(req))
	}
	return r
}
