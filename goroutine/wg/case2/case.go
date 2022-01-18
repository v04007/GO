package case2

import (
	"errors"
	"sync"
	"time"
)

func Case2(wg *sync.WaitGroup, ch *chan string, thiserr chan error) {
	time.Sleep(time.Second * 3)
	thiserr <- errors.New("执行case2 错误")
	close(*ch)
	wg.Done()
}
