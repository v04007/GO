package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}
type result struct {
	job *job
	sum int64
}

func ch1(jobChan chan<- *job) {
	wg.Done()
	for {
		newJob := &job{
			value: rand.Int63(),
		}
		jobChan <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func ch2(jobChan <-chan *job, resultChan chan<- *result) {
	wg.Done()
	for {
		jobs := <-jobChan
		n := jobs.value
		sum := int64(0)
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: jobs,
			sum: sum,
		}
		resultChan <- newResult
	}
}

var wg sync.WaitGroup
var a = make(chan *job, 100)
var b = make(chan *result, 100)

func main() {
	wg.Add(1)
	go ch1(a)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go ch2(a, b)
	}
	for v := range b {
		fmt.Printf("value %d sum %d \n", v.job, v.sum)
	}
	go wg.Wait()
}
