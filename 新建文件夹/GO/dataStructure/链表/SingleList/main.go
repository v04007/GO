package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickName string
	next     *HeroNode //表示指向下一个节点
}

// 给链表插入要给节点
// 第一种插入方法，直接在最后插入
func InsterHero(head *HeroNode, nextNode *HeroNode) {
	// 1、先找到链表最后节点
	// 2、创建一个辅助节点[跑龙套，协助]
	temp := head
	for {
		if temp.next == nil { //==nil 代表找到最后节点了
			break
		}
		temp = temp.next //让temp一直指向下一个temp节点
	}
	// 将nextNode加入到链表最后
	temp.next = nextNode
}

// 显示链表所有节点信息
func ListHeroNode(head *HeroNode) {
	// 1、创建一个辅助节点[跑龙套，协助]
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空")
		return
	}

	for {
		fmt.Printf("[%d %s %s]===>", temp.next.no, temp.next.name, temp.next.nickName)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}

}

func main() {

	head := &HeroNode{} //头节点采取默认即可
	// 创建第一个节点

	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickName: "及时雨",
	}

	InsterHero(head, hero1)
	ListHeroNode(head)
}
