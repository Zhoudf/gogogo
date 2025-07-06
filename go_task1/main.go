package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 只出现一次的数字：
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	// 创建一个 map 用于记录每个元素出现的次数
	countMap := make(map[int]int)

	// 遍历数组，统计每个元素出现的次数
	for _, num := range nums {
		countMap[num]++
	}

	// 遍历 map，找到出现次数为 1 的元素
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	// 若未找到，返回 0，实际题目保证有唯一解，不会走到这里
	return 0
}

// 判断一个整数是否是回文数
// 该方法通过反转整数的一半数字，然后与原数字的前半部分进行比较。
func isPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	// 将整数转换为字符串
	str := strconv.Itoa(x)
	reversedStr := ""
	// 倒序遍历字符串
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	// 比较原字符串和反转后的字符串是否相同
	return str == reversedStr
}

// 有效的括号
// 考察：字符串处理、栈的使用
// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
// isValid 判断括号字符串是否有效
func isValid(s string) bool {
	// 定义一个映射，存储右括号和对应的左括号
	mapping := map[rune]rune{
		')': '(', //key,value
		'}': '{', //key,value
		']': '[', //key,value
	}
	// 初始化一个栈，使用 rune 类型的切片
	stack := []rune{}

	// 遍历字符串中的每个字符
	for _, char := range s {
		switch char {
		// 遇到左括号，将其压入栈中
		case '(', '{', '[':
			stack = append(stack, char)
		// 遇到右括号
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}
	// 遍历结束后，如果栈为空，则字符串有效
	return len(stack) == 0
}

// 最长公共前缀,
// 考察：字符串处理、循环嵌套
// 题目：查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 以第一个字符串作为基准
	firstStr := strs[0]
	for i := 0; i < len(firstStr); i++ {
		char := firstStr[i]
		for j := 1; j < len(strs); j++ {
			// 如果当前字符串长度小于 i 或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != char {
				return firstStr[:i]
			}
		}
	}
	// 如果所有字符串都匹配第一个字符串的全部内容
	return firstStr
}

// 基本值类型
// 加一
// 难度：简单
// 考察：数组操作、进位处理
// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	n := len(digits)
	// 从数组的最后一位开始遍历
	for i := n - 1; i >= 0; i-- {
		// 当前位加 1
		digits[i]++
		// 如果当前位加 1 后小于 10，没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}
		// 若有进位，当前位设为 0，继续向前处理
		digits[i] = 0
	}
	// 如果所有位都产生了进位，需要在数组最前面添加一个 1
	result := make([]int, n+1)
	result[0] = 1
	return result
}

// 删除有序数组中的重复项，返回删除后数组的新长度
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 慢指针 i，初始指向第一个元素
	i := 0
	// 快指针 j，从第二个元素开始遍历
	for j := 1; j < len(nums); j++ {
		// 当 nums[i] 与 nums[j] 不相等时
		if nums[i] != nums[j] {
			// 将 nums[j] 赋值给 nums[i + 1]
			i++
			nums[i] = nums[j]
		}
	}
	// 返回不重复元素的长度
	return i + 1
}

// merge 合并所有重叠的区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用于存储合并后的区间
	result := [][]int{intervals[0]}

	// 遍历排序后的区间数组
	for _, interval := range intervals[1:] {
		last := result[len(result)-1]
		// 判断是否有重叠
		if interval[0] <= last[1] {
			// 有重叠，合并区间
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			// 没有重叠，添加新的区间
			result = append(result, interval)
		}
	}

	return result
}

func main() {
	//1.只出现一次的数字
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	fmt.Println(result)
	fmt.Println("-----------------")

	//2.回文数
	num := 12321
	fmt.Println(isPalindrome(num)) // 输出: true
	num = -12321
	fmt.Println(isPalindrome(num)) // 输出: false
	num = 10
	fmt.Println(isPalindrome(num)) // 输出: false
	fmt.Println("------------------")

	//3.有效的括号
	fmt.Println(isValid("()"))     // 输出: true
	fmt.Println(isValid("()[]{}")) // 输出: true
	fmt.Println(isValid("(]"))     // 输出: false
	fmt.Println(isValid("([)]"))   // 输出: false
	fmt.Println(isValid("{[]}"))   // 输出: true
	fmt.Println("------------------")

	//4.最长公共前缀
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1)) // 输出: "fl"
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs2)) // 输出: ""
	strs3 := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs3)) // 输出: "a"
	fmt.Println("------------------")

	//5.加一
	digits1 := []int{1, 2, 3}
	fmt.Println(plusOne(digits1)) // 输出: [1 2 4]
	digits2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits2)) // 输出: [4 3 2 2]
	digits3 := []int{9, 9, 9}
	fmt.Println(plusOne(digits3)) // 输出: [1 0 0 0]
	fmt.Println("------------------")

	//6.删除有序数组中的重复项
	nums := []int{1, 1, 2, 2, 3, 4, 4, 4, 5}
	newLength := removeDuplicates(nums)
	fmt.Println("新长度:", newLength)
	fmt.Println("去重后的数组:", nums[:newLength])
	fmt.Println("------------------")

	//7.合并区间
	intervals := [][]int{
		{1, 3},
		{8, 10},
		{2, 6},
		{15, 18},
	}
	fmt.Println(merge(intervals)) // 输出: [[1 6] [8 10] [15 18]]

	intervals2 := [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(intervals2)) // 输出: [[1 5]]
}
