package engine

import (
	"log"
	"spider/crawler/fetcher"
)

type SimpleEngine struct {
}

func worker(r Request) (ParseResult, error) {
	if r.Url == "" {
		return ParseResult{}, nil
	}
	body, err := fetcher.Fetch(r.Url)
	//log.Printf("%s\n", r.Url)
	if err != nil {
		log.Printf("Fetcher:error"+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	exist := make(map[string]bool)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		_, ok := exist[r.Url]
		if ok {
			continue
		}
		if r.Url == "" {
			//log.Println("go to stroe")
			continue
		}
		exist[r.Url] = true

		parseResult1, err := worker(r)
		if err != nil {
			log.Printf("Fetcher:error"+"fetching url %s: %v", r.Url, err)
			continue
		}
		requests = append(requests, parseResult1.Requests...)
		for _, item := range parseResult1.Items {
			log.Printf("Got item %v", item)
		}
	}
}
