package str

import (
	"errors"
	"fmt"
	"os"
)

// 链表不能有效利用数组
// 队列遵循先入先出原则

type CircleQueue struct {
	maxSize int    //队列最大容量
	array   [5]int //数组=>模拟队列
	head    int    //表示指向队列 队首随着输出而改变
	tail    int    //表示指向队列 尾部随着输入而改变
}

// 入队列
func (this *CircleQueue) Push(val int) (err error) {
	//判断队列是否已满
	if this.IsFull() {
		return errors.New("queue full")
	}
	//分析出this.tail 在队列尾部，但是包含最后的元素
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队列
func (this *CircleQueue) Pop() (val int, err error) {
	// 先判断队列是否为空
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// IsFull 判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// IsEmpty 判断环形队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// Size 取出环形队列有多少个元素，关键的算法
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

// ListQueue 显示队列,找到队首,然后到遍历到队尾
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下:")
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	// 设计一个辅助的变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

//编写一个主函数测试,测试
func mains() {
	//先创建一个队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
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
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
