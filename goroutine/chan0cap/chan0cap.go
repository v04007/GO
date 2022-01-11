package main

import (
	"fmt"
	"time"
)

func chancap() {
	sedingInterval := time.Second
	receptionInterval := time.Second * 2
	intChan := make(chan int, 0)
	go func() {
		var ts0, ts1 int64
		for i := 0; i <= 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent:%d [interval:%d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sedingInterval)
		}
		close(intChan)
	}()
	var ts0, ts1 int64
Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Received:", v)
			} else {
				fmt.Printf("Received:%d [interval: %d s]\n", v, ts1-ts0)
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}
	fmt.Println("End.")
}

func timerbase() {
	timer := time.NewTimer(2 * time.Second) //在多少秒后发送当前世界
	fmt.Printf("Present time:%v.\n", time.Now())
	expirationTime := <-timer.C //会一直阻塞直到定时器到期
	fmt.Printf("Expiration time:%v.\n", expirationTime)
	fmt.Printf("Stop timer:%v.\n", timer.Stop())
}

func chantimeout() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1 //一秒后写入
	}()
	select {
	case e := <-intChan:
		fmt.Printf("Received:%v", e)
	case <-time.NewTimer(time.Millisecond * 500).C: //500ms打印超时
		fmt.Println("Timeout!")
	}
}
func main() {
	//chancap()
	//timerbase()
	//chantimeout()
}
