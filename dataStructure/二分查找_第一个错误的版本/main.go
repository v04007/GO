package main

import "fmt"

func firstBadVersion(n int) int {
	left, right := 1, n //left 从第一个开始计算递增，right最新版本号
	for left <= right { //left 递增不能大于最新版本
		mid := (right-left)>>1 + left //中位数
		if isBadVersion(mid) {        //判断中位数是不是第一个错误版本
			right = mid - 1 //最大版本数，由中位数向左边移动一位
		} else if mid < n { //中位数小于最新版本号
			left = mid + 1 //开始位等于中位数向右边前进一位
		}
	}
	return left
}

func isBadVersion(n int) bool { //第一个错误版本
	return n == 3
}

func main() {
	version := firstBadVersion(5)
	fmt.Println("version", version)
}
