package main

import (
	"errors"
	"fmt"
	"os"
)

//队列是一个有序列表，可以用数组或是链表来实现。  遵循先入先出的原则。即：先存入队列的数据，要先取出。后存入的要后取出
//单向链表不能有效利用数组

type Queue struct {
	maxSize int    //队列最大容量
	array   [5]int //数组=>模拟队列
	front   int    //表示指向队列队首随着输出而改变
	rear    int    //表示指向队列尾部随着输入而改变
}

func (q *Queue) AddQueue(val int) (err error) {
	//判断队列是否已满
	if q.rear == q.maxSize-1 { //重要提示:rear 是队列尾部(含最后元素)
		return errors.New("queue full")
	}
	q.rear++ //rear 后移
	q.array[q.rear] = val
	return
}

//显示队列,找到队首,然后到遍历到队尾
func (q *Queue) shouQueue() {
	fmt.Println("队列当前情况是:")
	//q.front 不包含队首的元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
}

func (q *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if q.rear == q.front { //队空
		return -1, errors.New("queue empty")
	}
	q.front++
	val = q.array[q.front]
	return val, err
}

//编写一个主函数测试,测试
func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1、输入add 表示添加数据到队列")
		fmt.Println("2、输入get 表示从队列获取数据")
		fmt.Println("3、输入show 表示显示队列")
		fmt.Println("4、输入exit 表示退出队列")

		fmt.Scan(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数")
			fmt.Scan(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.shouQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
