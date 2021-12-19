package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"thread-pool/api/words"
	"time"
)

var numberOfWorkers int
var numberOfRandomWords int
var waitGroup sync.WaitGroup

/* Initialize the amount of workers according to the number of CPU existent on the running machine. */
func init() {
	numberOfWorkers = runtime.NumCPU()
	numberOfRandomWords = 10000
}

func readMessage(queueChannel chan string, workerId int) {
	for msg := range queueChannel {
		fmt.Printf("Worker %d Got the Message: %s\n", workerId, msg)
	}
	waitGroup.Done()
}

func main() {
	/* queueChannel is the pool which all the workers are listening to. All the data processed sent to the queue Channel
	will be picked up by an available worker. */
	queueChannel := make(chan string, 250)
	launchWorkers(queueChannel, numberOfWorkers)
	waitGroup.Add(numberOfWorkers)
	startTime := time.Now()

	// Send values to queue
	wordsApi := words.NewApi("https://random-word-api.herokuapp.com/word")
	randomWords, err := wordsApi.GetRandom(numberOfRandomWords)
	if err != nil {
		log.Fatal(err)
	}
	for _, word := range randomWords {
		queueChannel <- word
	}

	close(queueChannel)
	timeElapsed := time.Since(startTime)
	waitGroup.Wait()
	fmt.Printf("Process took %s seconds.\n", timeElapsed)
}

func launchWorkers(queueChannel chan string, workersAmount int) {
	for workerId := 0; workerId < workersAmount; workerId++ {
		go readMessage(queueChannel, workerId)
	}
}
