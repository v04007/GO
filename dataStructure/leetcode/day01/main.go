package main

//twoSum LeetCode第一题
//相对双层for循环更快，巧用map
//map中key为nums的value，map中的value为nums的key
//通过相减后的差做下做key，取val，没有对应val就添加进去
//如果有值就直接取出来，不用多层for嵌套
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

//twoSum LeetCode第二题
func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	//(x%10 == 0 && x != 0) 则其第一位数字也应该是 0 只有 0 满足这一属性，其他如10、20都不满住，return false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10 //第一次执行revertedNumber为0，0*10为0加上 x%10的余数，依次类推，就将数倒过来了
		x /= 10                                   //x为整型没有小数
	}
	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

//romanToInt LeetCode第三题
func romanToInt(s string) (ans int) {
	symbolValues := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		//i < n-1当为倒数第二个时就不走这个逻辑，防止下标越界，
		//因为value < symbolValues[s[i+1]]中的i+1执行到最后一个时还加1就会报panic下标越界
		//当前值小于下一个就代表当前值为负数，例如IV（-5），I小于V（1小于5）
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}

func main() {
	//nums := []int{2, 7, 11, 15}
	//target := 9
	//fmt.Println(twoSum(nums, target))

	//fmt.Println(isPalindrome(121))

	//fmt.Println(romanToInt("IX"))

}
