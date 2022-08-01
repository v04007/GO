package main

import "fmt"

func Rescuvie(n int) (result int) {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

//{5 * Rescuvie(4)}
//{5 * {4 * Rescuvie(3)}}
//{5 * {4 * {3 * Rescuvie(2)}}}
//{5 * {4 * {3 * {2 * Rescuvie(1)}}}}
//{5 * {4 * {3 * {2 * 1}}}}
//{5 * {4 * {3 * 2}}}
//{5 * {4 * 6}}
//{5 * 24}
//120

func sum(n int) int {
	total := 0
	// 从1加到N, 1+2+3+4+5+..+N
	for i := 1; i <= n; i++ {
		total = total + i
	}
	return total
}

func sum2(n int) int {
	total := ((1 + n) * n) / 2
	return total
}

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func middleNode(head *LinkNode) *LinkNode {
	Fast, slow := head, head
	for Fast != nil && Fast.NextNode != nil {
		Fast = Fast.NextNode.NextNode //Fast 是 slow 的两倍
		slow = slow.NextNode
	}
	return slow
}

//func removeNthFromEnd(head *LinkNode, n int) *LinkNode {
//
//	// 方便删除头节点
//	var dummy = &LinkNode{NextNode: head}
//
//	slow, fast := dummy, dummy
//	for i := 0; i < n; i++ {
//		fast = fast.NextNode
//	}
//
//	for fast.NextNode != nil {
//		slow, fast = slow.NextNode, fast.NextNode
//	}
//
//	slow.NextNode = slow.NextNode.NextNode //删除
//	return dummy.NextNode
//}

func removeNthFromEnd(head *LinkNode, n int) *LinkNode {
	dummy := new(LinkNode)
	dummy.NextNode = head
	quick, slow := dummy, dummy

	for i := 0; quick.NextNode != nil && i < n; i++ {
		quick = quick.NextNode
	}

	for quick.NextNode != nil && slow.NextNode != nil {
		quick = quick.NextNode
		slow = slow.NextNode
	}
	slow.NextNode = slow.NextNode.NextNode
	return dummy.NextNode
}

func linkedList() {
	node := new(LinkNode)
	node.Data = 1

	node1 := new(LinkNode)
	node1.Data = 2
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 3
	node1.NextNode = node2

	node3 := new(LinkNode)
	node3.Data = 4
	node2.NextNode = node3

	node4 := new(LinkNode)
	node4.Data = 5
	node3.NextNode = node4

	node5 := new(LinkNode)
	node5.Data = 6
	node4.NextNode = node5

	node6 := new(LinkNode)
	node6.Data = 7
	node5.NextNode = node6

	//nowNode := node
	////for {
	////	if nowNode != nil {
	////		fmt.Println(nowNode.Data)
	////		nowNode = nowNode.NextNode
	////		continue
	////	}
	////	break
	////}

	//linkNode := middleNode(node)
	end := removeNthFromEnd(node, 2)
	fmt.Println(
		end)
	//end.NextNode,
	//end.NextNode.NextNode,
	//end.NextNode.NextNode.NextNode,
	//end.NextNode.NextNode.NextNode.NextNode,
	//end.NextNode.NextNode.NextNode.NextNode.NextNode,
	//end.NextNode.NextNode.NextNode.NextNode.NextNode.NextNode)
}

func main() {
	//fmt.Println(sum(10))
	linkedList()
}
