package main

import "fmt"
func binaryserrch(nums []int, target int) int {
	left, right := 0, len(nums)-1 //开始下表为0，结束下标为传入长度-1
	for left <= right {           //for 循环结束条件为left小于right
		mid := (right + left) / 2 //中位数为开始+传入长度下表 除以2
		if nums[mid] == target {  //中位数下标对应的值等于目标值，人会中位数
			return mid
		} else if nums[mid] > target { //中位数下标值大于目标值
			left = mid - 1 //开始为向左边移动一位
		} else if nums[mid] < target { //中位数下标值小于目标值
			left = mid + 1 //开始为向右边移动一位
		}
	}
	return -1
}

func main()  {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := binaryserrch(list, 9)
	fmt.Println("target", target)
}