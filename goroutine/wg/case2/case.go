package case2

import (
	"sync"
	"time"
)

func Case2(wg *sync.WaitGroup, ch *chan string) {
	time.Sleep(time.Second * 3)
	*ch <- "执行case2"
	wg.Done()
}
