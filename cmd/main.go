package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"runtime"
)

const TotalIterations = int64(100_000_000)

func computeRange(start, end int64) int64 {
	result := int64(0)
	for i := start; i < end; i++ {
		val := (i * 31) ^ (i + 17)
		result ^= val
	}
	return result
}

func singleThreaded() {
	start := time.Now()
	result := computeRange(0, TotalIterations)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Single-threaded: %.6f\n", float64(elapsed.Microseconds())/float64(1_000_000))
	fmt.Printf("Result: %d\n", result)
}

func multiThreaded(concurrency int64) {
	// divide total iterations to chunks
	chunkSize := TotalIterations / concurrency

	resultChan := make(chan int64, concurrency)

	var wg sync.WaitGroup
	wg.Add(int(concurrency))

	startTime := time.Now()

	for i := int64(0); i < concurrency; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == concurrency-1 {
			end = TotalIterations
		}
		go func(start, end int64) {
			defer wg.Done()
			resultChan <- computeRange(start, end)
		}(start, end)
	}
	wg.Wait()

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)

	result := int64(0)
	for i := int64(0); i < concurrency; i++ {
		result ^= <-resultChan
	}

	fmt.Printf("Concurrent: %.6f seconds\n", float64(elapsed.Microseconds())/float64(1_000_000))
	fmt.Printf("Result: %d\n", result)
}

func sysInfo() error {

	return nil
}

func main() {
	// command-line flags
	concurrency := flag.Int64("concurrency", int64(runtime.NumCPU()), "Number of 'threads' to use")
	flag.Parse()

    // system info
	if err := sysInfo(); err != nil {
		log.Printf("[cpuinfo]: failed to get cpuinfo: %v", err)
	}

    // custom concurrency or not?
	if *concurrency != int64(runtime.NumCPU()) {
		fmt.Printf("[tinybench]: setting custom concurrency to %d", *concurrency)
	} else {
		fmt.Printf("[tinybench]: using default CPU concurrency of %d", *concurrency)
	}

	fmt.Printf("Running compute benchmark with %d iterations...\n", TotalIterations)
	singleThreaded()
	fmt.Printf("Running multithreaded compute benchmark with %d iterations on concurrency %d...\n", TotalIterations, *concurrency)
	multiThreaded(*concurrency)
}
