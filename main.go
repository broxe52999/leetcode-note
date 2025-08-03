package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

// 轮转数组
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// H指数
func hIndex(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		if citations[i] >= n-i {
			ans = n - i
		}
	}
	return ans
}

// 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return []int{0}
	}
	answer := make([]int, n)
	answer[0] = nums[0]
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i]
	}
	temp := 1
	for i := n - 1; i >= 1; i-- {
		if i != n-1 {
			temp = temp * nums[i+1]
		}
		answer[i] = answer[i-1] * temp
	}
	answer[0] = temp
	return answer
}
