package engine

import (
	"log"
	"spider/crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler        Schduler
	WorkerCount      int
	ItemSaver        chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

func Worker(r Request) (ParseResult, error) {
	if r.Url == "" {
		return ParseResult{}, nil
	}
	body, err := fetcher.Fetch(r.Url)
	//log.Printf("%s\n", r.Url)
	if err != nil {
		log.Printf("Fetcher:error"+"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}

type Schduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if url == "" {
		return false
	}
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request :"+"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func(i Item) {
				e.ItemSaver <- i
			}(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request :" + "%s" ,request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}
