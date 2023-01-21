package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"infowatchtest/fstat"
)

func main() {
	filesDir, err := filepath.Abs("./input")
	if err != nil {
		log.Fatalf("filepath.Abs err: %s", err)
	}

	pool := fstat.NewWorkerPool(runtime.NumCPU())

	// Getting ready to proceed jobs
	pool.Start()

	// Adding Jobs async
	go addJobs(pool, filesDir)

	// Printing result
	for res := range pool.ResC() {
		if res.Err != nil {
			fmt.Printf("Histogram %s err: %s\n\n", res.ID, res.Err.Error())

			continue
		}

		fmt.Println("Histogram " + res.ID)
		res.Hist.PrintSorted()
		fmt.Println()
	}
}

func addJobs(pool *fstat.WorkerPool, filesDir string) {
	defer pool.Finish()

	// Reading files
	files, err := os.ReadDir(filesDir)
	if err != nil {
		panic(fmt.Sprintf("readDir err: %s", err))
	}

	for _, file := range files {
		// skipping directories
		if file.IsDir() {
			continue
		}

		pool.AddJob(&fstat.Job{
			ID:    file.Name(),
			FPath: filepath.Join(filesDir, file.Name()),
		})
	}
}
