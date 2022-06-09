package main

import "fmt"

type HeroNode struct {
	no             int
	name, nickName string
	next           *HeroNode
}

func InsterHero(head *HeroNode, nextNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = nextNode
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]====>", temp.no, temp.name, temp.nickName)
		temp := temp.next
	}
}

func main() {

}
