package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"infowatchtest/fstat"
)

func main() {
	// Result histogram of each file
	resC := make(chan *fstat.Result)

	// Files to proceed
	jobC := make(chan *fstat.Job)

	// Flag, when all workers are finished
	wg := sync.WaitGroup{}

	// Fan-out workers limited to CPU number
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)

		go func() {
			fstat.WorkerStart(jobC, resC)
			wg.Done()
		}()
	}

	// Close result channel when all workers finished its job
	go func() {
		wg.Wait()
		close(resC)
	}()

	// Async adding new jobs to job's channel
	go func() {
		filesDir, err := filepath.Abs("./input")
		if err != nil {
			log.Fatalf("filepath.Abs err: %s", err)
		}

		// Reading files
		files, err := os.ReadDir(filesDir)
		if err != nil {
			log.Fatalf("readDir err: %s", err)
		}

		for _, file := range files {
			// skipping directories
			if file.IsDir() {
				continue
			}

			jobC <- &fstat.Job{
				ID:    file.Name(),
				FPath: filepath.Join(filesDir, file.Name()),
			}
		}

		close(jobC)
	}()

	// Printing result
	for res := range resC {
		if res.Err != nil {
			fmt.Printf("Histogram %s err: %s\n\n", res.ID, res.Err.Error())

			continue
		}

		fmt.Println("Histogram " + res.ID)
		res.Hist.PrintSorted()
		fmt.Println()
	}
}
