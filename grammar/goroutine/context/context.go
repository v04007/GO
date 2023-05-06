package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

//var wg sync.WaitGroup
//
//func gen(ctx context.Context) <-chan int {
//	dst := make(chan int)
//	n := 1
//	go func() {
//		for {
//			select {
//			case <-ctx.Done():
//				return
//			case dst <- n:
//				n++
//			}
//		}
//	}()
//	return dst
//}
//func WithCancel() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	for n := range gen(ctx) {
//		fmt.Println(n)
//		if n == 5 {
//			break
//		}
//	}
//}
//
//func WithDeadline() {
//	d := time.Now().Add(time.Millisecond * 1000)
//	fmt.Println(d)
//	ctx, cancel := context.WithDeadline(context.Background(), d)
//	defer cancel()
//
//	select {
//	case <-ctx.Done():
//		fmt.Println(ctx.Err())
//	case <-time.After(2 * time.Second):
//		fmt.Println("执行了")
//	}
//}
//
//func worker(ctx context.Context) {
//LOOP:
//	for {
//		fmt.Println("connenttime ...")
//		time.Sleep(time.Millisecond * 10)
//		select {
//		case <-ctx.Done():
//			break LOOP
//		default:
//		}
//	}
//	fmt.Println("worker ...")
//	wg.Done()
//}
//func WithTimeout() {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 5)
//	cancel()
//	wg.Wait()
//	fmt.Println("over")
//}
//
//type TractCood string
//
//func valWorker(ctx context.Context) {
//	key := TractCood("TRACE_CODE")
//	tractcood, ok := ctx.Value(key).(string)
//	if !ok {
//		fmt.Println("invalid tract cood")
//	}
//LOOP:
//	for {
//		fmt.Println("tract cood:", tractcood)
//		time.Sleep(time.Millisecond * 10)
//		select {
//		case <-ctx.Done():
//			break LOOP
//		default:
//		}
//	}
//	fmt.Println("worker done")
//	wg.Done()
//}
//
//func WithValue() {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
//	ctx = context.WithValue(ctx, TractCood("TRACE_CODE"), "20210629")
//	wg.Add(1)
//	go valWorker(ctx)
//	time.Sleep(time.Second * 5)
//	cancel()
//	wg.Wait()
//	fmt.Println("over")
//}

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

//func f1(ctx context.Context) {
//	defer wg.Done()
//	go f2(ctx)
//FARLOOK:
//	for {
//		time.Sleep(time.Millisecond * 500)
//		fmt.Println("f1:", time.Now().Format("05"))
//		select {
//		case <-ctx.Done():
//			break FARLOOK
//		default:
//		}
//	}
//}

//func main() {
//	cherrors := make(chan error)
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go func() {
//		//cherrors <- errors.New("出错啦。。我是错误信息")
//		wg.Done()
//	}()
//	go func() {
//		time.Sleep(time.Second * 2)
//		cherrors <- errors.New("出错啦。。我是错误信息")
//		wg.Done()
//	}()
//
//	go func() {
//		for {
//			select {
//			case err := <-cherrors:
//				close(cherrors)
//				fmt.Println(err)
//			}
//		}
//	}()
//
//	wg.Wait()
//
//	fmt.Println("结束")
//}

func main() {
	group := new(errgroup.Group)

	nums := []int{-1, 0, 1}
	for _, num := range nums {
		num := num
		group.Go(func() error {
			res, err := output(num)
			fmt.Println(res)
			return err
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}
}

func output(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("math: square root error!")
	}
	return num, nil
}
