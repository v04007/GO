package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func WithDeadline() {
	d := time.Now().Add(time.Millisecond * 1000)
	fmt.Println(d)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(2 * time.Second):
		fmt.Println("执行了")
	}
}

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("connenttime ...")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker ...")
	wg.Done()
}
func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

type TractCood string

func valWorker(ctx context.Context) {
	key := TractCood("TRACE_CODE")
	tractcood, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid tract cood")
	}
LOOP:
	for {
		fmt.Println("tract cood:", tractcood)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}

func WithValue() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TractCood("TRACE_CODE"), "20210629")
	wg.Add(1)
	go valWorker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

// 使用自定义
// var exitChan = make(chan bool, 1)

// func f1() {
// 	defer wg.Done()
// FARLOOK:
// 	for {
// 		fmt.Println(time.Now().Format("05"))
// 		time.Sleep(time.Millisecond * 500)
// 		select {
// 		case <-exitChan:
// 			break FARLOOK
// 		default:
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go f1()
// 	time.Sleep(time.Second * 5)
// 	exitChan <- true
// 	wg.Wait()
// }

func f2(ctx context.Context) {
	defer wg.Done()
FARLOOK:
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("f2:", time.Now().Format("05"))
		select {
		case <-ctx.Done():
			break FARLOOK
		default:
		}
	}
}

func f1(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FARLOOK:
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("f1:", time.Now().Format("05"))
		select {
		case <-ctx.Done():
			break FARLOOK
		default:
		}
	}
}

func sel(g coroutine, i int) {
	fmt.Println(i)
	if i == 5 {
		g.err <- struct{}{}
		g.wg.Done()
	}
	g.wg.Done()
}

type coroutine struct {
	wg  *sync.WaitGroup
	ctx context.Context
	err chan struct{}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //使用ctx结束goroutine

	routine := coroutine{
		ctx: ctx,
		wg:  &sync.WaitGroup{},
		err: make(chan struct{}),
	}
	for i := 0; i < 10; i++ {
		routine.wg.Add(1)
		go sel(routine, i)
	}

	select {
	case <-routine.err:
		cancel()
	}
	wg.Wait()
}
