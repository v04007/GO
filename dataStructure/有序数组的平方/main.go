package main

import "fmt"

func sortedSquares(nums []int) []int {
	list := make([]int, len(nums), len(nums))
	left, right, ind := 0, len(nums)-1, len(nums)-1 //left头部从0开始，right尾部len(nums)-1结束，ind等于len(nums)-1结
	for left <= right && ind >= 0 {
		if nums[left]*nums[left] > nums[right]*nums[right] { //头部大于尾部
			list[ind] = nums[left] * nums[left] //ind位值等于头部数平方
			left += 1                           //left向右移动一位
		} else {
			list[ind] = nums[right] * nums[right] //ind位值等于头部数平方
			right -= 1                            //right向右移动一位
		}
		ind -= 1 //减一
	}
	return list
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) //递增顺序
}
