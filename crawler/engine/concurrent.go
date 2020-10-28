package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Schduler
	WorkerCount int
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

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
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
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request :"+"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	count := 1
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got %d item: %v", count, item)
			count++
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
