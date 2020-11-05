package engine

import (
	"log"
)

type SimpleEngine struct {
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

		parseResult1, err := Worker(r)
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
