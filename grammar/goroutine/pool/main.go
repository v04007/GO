package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func busi(ch chan struct{}, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
	wg.Done()
}

func main() {
	//模拟用户需求go业务的数量
	task_cnt := math.MaxInt64
	ch := make(chan struct{}, 3)
	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		ch <- struct {
		}{}
		go busi(ch, i)
	}

	wg.Wait()
}
