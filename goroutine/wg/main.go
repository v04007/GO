package main

import (
	"fmt"
	"sync"
	"time"
	"wg/case1"
	"wg/case2"
)

type info struct {
	name1 string
	name2 string
}

func main() {
	now := time.Now()
	str := new(info)
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)
	wg.Add(2)
	go case1.Case1(&wg, &ch1)
	go case2.Case2(&wg, &ch2)

	if c, ok := <-ch1; ok {
		str.name2 = c
	}
	if c, ok := <-ch2; ok {
		str.name1 = c
	}

	wg.Wait()
	fmt.Println("main执行完毕", str)
	fmt.Println(now.Sub(time.Now()))
}
