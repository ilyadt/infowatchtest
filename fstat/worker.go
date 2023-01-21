package fstat

import (
	"infowatchtest/histogram"
)

type Job struct {
	ID    string
	FPath string
}

type Result struct {
	ID   string
	Err  error
	Hist *histogram.DiscreteHistogram[byte]
}

func WorkerStart(jobC <-chan *Job, resC chan<- *Result) {
	for job := range jobC {
		jobFpath := job.FPath
		jobID := job.ID

		hist, err := GatherStats(jobFpath)

		resC <- &Result{
			ID:   jobID,
			Err:  err,
			Hist: hist,
		}
	}
}
