package scheduler

import "spider/crawler/engine"

type QueuedScheduler struct {
	requetsChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requetsChan <- request
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	q.workerChan = make(chan chan engine.Request)
	q.requetsChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var vcworkerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(vcworkerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = vcworkerQ[0]
			}
			select {
			case r := <-q.requetsChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				vcworkerQ = append(vcworkerQ, w)
			case activeWorker <- activeRequest:
				vcworkerQ = vcworkerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()
}
