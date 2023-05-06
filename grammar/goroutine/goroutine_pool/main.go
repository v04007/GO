package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

//自定义协程池
func busi(ch chan struct{}, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
	<-ch
	wg.Done()
}
func Pool(amount int64) {
	//模拟用户需求go业务的数量
	task_cnt := math.MaxInt64
	pool := make(chan struct{}, 3)
	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		pool <- struct {
		}{}
		go busi(pool, i)
	}
	wg.Wait()
}

func main() {
	Pool(20)
}
