package case1

import (
	"sync"
	"time"
)

func Case1(wg *sync.WaitGroup, ch *chan string) {
	time.Sleep(time.Second * 5)
	*ch <- "执行完毕case1"
	close(*ch)
	wg.Done()
}
