package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type CircleQueue struct {
	array               [5]int
	maxSize, head, tail int
}

func (c *CircleQueue) Push(arg int) {
	if c.IsFull() {
		log.Fatal("队列已满")
		return
	}
	c.array[c.tail] = arg
	c.tail = (c.tail + 1) % c.maxSize
}

func (c *CircleQueue) IsEmpty() bool {
	return c.tail+1 == c.head
}

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return -1, errors.New("队列为空")
	}
	val = c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return
}
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}
func (c *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下:")
	size := c.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	// for i := 0; i < size; i++ {
	// 	fmt.Printf("arr[%d]=%d\t", c.head, c.array[c.head])
	// 	c.head = (c.head + 1) % c.maxSize
	// }

	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}

	fmt.Println()
}
func main() {
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	for {
		fmt.Println("输入操作指令")
		var key string
		var val int
		fmt.Scan(&key)
		switch key {
		case "add":
			fmt.Println("输入存入值：")
			fmt.Scan(&val)
			queue.Push(val)
		case "get":
			val2, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("取出值为")
			fmt.Println(val2)
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(1)
		}
	}
}
