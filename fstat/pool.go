package fstat

import (
	"sync"

	"infowatchtest/histogram"
)

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		jobC: make(chan *Job),
		resC: make(chan *Result),
		sem:  make(chan struct{}, numWorkers),
	}
}

type WorkerPool struct {
	jobC chan *Job
	resC chan *Result
	sem  chan struct{} // semaphore
}

type Job struct {
	ID    string
	FPath string
}

type Result struct {
	ID   string
	Err  error
	Hist *histogram.DiscreteHistogram[rune]
}

func (p *WorkerPool) AddJob(j *Job) {
	p.jobC <- j
}

func (p *WorkerPool) ResC() <-chan *Result {
	return p.resC
}

func (p *WorkerPool) Start() {
	go start(p.sem, p.jobC, p.resC)
}

func (p *WorkerPool) Finish() {
	close(p.jobC)
}

func start(sem chan struct{}, jobC <-chan *Job, resC chan<- *Result) {
	wg := sync.WaitGroup{}

	for job := range jobC {
		wg.Add(1)
		sem <- struct{}{}

		jobFpath := job.FPath
		jobID := job.ID

		go func() {
			defer func() {
				<-sem
				wg.Done()
			}()

			hist, err := GatherStats(jobFpath)

			resC <- &Result{
				ID:   jobID,
				Err:  err,
				Hist: hist,
			}
		}()
	}

	wg.Wait()

	close(resC)
}
