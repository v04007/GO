package main

import (
	"fmt"
)

//案例1	chan
//var strChan = make(chan string, 3)
//
//func main() {
//	syncbase1 := make(chan struct{}, 1)
//	syncbase2 := make(chan struct{}, 2)
//
//	go func() {
//		//阻塞，当接受到信号后执行
//		<-syncbase1
//		fmt.Println("Received a sync and wait a second ... [receiver]")
//		time.Sleep(time.Second)
//		for {
//			if elem, ok := <-strChan; ok {
//				fmt.Println("Received:", elem, "[]receiver")
//			} else {
//				break
//			}
//		}
//		fmt.Println("Stopped. [receiver]")
//		syncbase2 <- struct{}{}
//	}()
//
//	go func() {
//		for _, elem := range []string{"a", "b", "c", "d"} {
//			strChan <- elem
//			fmt.Println("Sent", elem, "[sender]")
//			if elem == "c" {
//				//当为c向syncbase1 chan发送信号,
//				syncbase1 <- struct{}{}
//				fmt.Println("Sent a sync signal. [sender]")
//			}
//		}
//		fmt.Println("Sent a sync signal ...[sender]")
//		time.Sleep(time.Second * 2)
//		close(strChan)
//		syncbase2 <- struct{}{}
//	}()
//
//	<-syncbase2
//	<-syncbase2
//}

//案例2	chan传递的值是拷贝值
//var mapChan = make(chan map[string]int, 1)
//
//func main() {
//	syncChan := make(chan struct{}, 2)
//	go func() {
//		for {
//			//1 先阻塞
//			if elem, ok := <-mapChan; ok {
//				//3 对传入的拷贝值make(map[string]int)执行++,0,1,2,3,4
//				fmt.Println("传入值为", elem["count"])
//				elem["count"]++
//			} else {
//				break
//			}
//		}
//		fmt.Println("Stopped [receiver]")
//		syncChan <- struct{}{}
//	}()
//
//	go func() {
//		countMap := make(map[string]int)
//		for i := 0; i < 5; i++ {
//			//2 写入信号
//			mapChan <- countMap
//			time.Sleep(time.Millisecond)
//			fmt.Println("The count map: %v .[sender]", countMap)
//		}
//		close(mapChan)
//		syncChan <- struct{}{}
//	}()
//
//	<-syncChan //阻塞main
//	<-syncChan //阻塞main
//}

//案例三  传递切片类型的值

//type Counter struct {
//	count int
//}
//
//var mapChan = make(chan map[string]Counter, 1) //打印0值
//
////var mapChan = make(chan map[string]*Counter, 1) 为这种类型打印的是地址
//
//func main() {
//	syncChan := make(chan struct{}, 2)
//	go func() {
//		for {
//			//1 先阻塞
//			if elem, ok := <-mapChan; ok {
//				counter := elem["count"]
//				counter.count++
//			} else {
//				break
//			}
//		}
//		fmt.Println("Stopped [receiver]")
//		syncChan <- struct{}{}
//	}()
//
//	go func() {
//		countMap := map[string]Counter{"count": Counter{}}
//		//countMap := map[string]*Counter{"count": &Counter{}} 打印地址
//		for i := 0; i < 5; i++ {
//			mapChan <- countMap
//			time.Sleep(time.Millisecond)
//			fmt.Printf("The count map: %v .[sender]\n", countMap)
//		}
//		close(mapChan)
//		syncChan <- struct{}{}
//	}()
//
//	<-syncChan //阻塞main
//	<-syncChan //阻塞main
//}

//案例四 工具表达式第二个解决判断通道是否已关闭和还有没有值可取

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() { //用于演示接受操作
		<-syncChan1 //1 先阻塞
		for {
			if elem, ok := <-dataChan; ok { //5 获取chan中写入的5个i值，to
				fmt.Printf("Received:%d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Stopped [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i //2 写入i值
			fmt.Printf("Sent: %v .[sender]\n", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}       //3 发送信号
		fmt.Println("Done. [sender]") //4 该协程执行完毕
		syncChan2 <- struct{}{}
	}()

	<-syncChan2 //阻塞main
	<-syncChan2 //阻塞main
}
