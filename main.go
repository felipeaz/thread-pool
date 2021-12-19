package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var numberOfWorkers int
var waitGrup sync.WaitGroup

/* Initialize the amount of workers according to the number of CPU existent on the running machine. */
func init() {
	numberOfWorkers = runtime.NumCPU()
}

func calculateArea(queueChannel <-chan string, workerId int) {
	for msg := range <-queueChannel {
		fmt.Printf("Worker %d Got the Message: %s", workerId, msg)
	}
	waitGrup.Done()
}

func main() {
	/* queueChannel is the pool which all the workers are listening to. All the data processed sent to the queue Channel
	will be picked up by an available worker. */
	queueChannel := make(chan string, 250)
	launchWorkers(queueChannel, numberOfWorkers)
	waitGrup.Add(numberOfWorkers)
	startTime := time.Now()
	// Send values to queue
	close(queueChannel)
	timeElapsed := time.Since(startTime)
	waitGrup.Wait()
	fmt.Printf("Process took %s seconds.\n", timeElapsed)
}

func launchWorkers(queueChannel <-chan string, workersAmount int) {
	for workerId := 0; workerId < workersAmount; workerId++ {
		go calculateArea(queueChannel, workerId)
	}
}
