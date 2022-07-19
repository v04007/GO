package main

import "fmt"

func searchInsert(nums []int, target int) int {
	i := len(nums)        //i 初始值为nums长度
	idx := i              //下标为i值
	left, right := 0, i-1 //left从0开始，从nums长度-1结束
	for left <= right {   //结束条件为left(开始位)小于等于right(结束位)
		mid := (right-left)>>1 + left //计算中位数
		if nums[mid] >= target {      //中位数小于等于目标值
			idx = mid       //记录下中位数值
			right = mid - 1 //结束值等于中位数向左边移动一位
		} else {
			left = mid + 1 //开始值向中位数向右边移动一位
		}
	}
	return idx
}

func main() {
	insert := searchInsert([]int{1, 3, 5, 6}, 5) //输出2
	fmt.Println(insert)
}
