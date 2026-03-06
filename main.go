package main

import (
	"fmt"
	"sort"
)

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

func makeAndNew() {
	p := new(int)
	fmt.Println(*p)
	*p = 10
	fmt.Println(*p)

	s := make([]int, 5, 10)
	fmt.Println(s)
	m := make(map[string]int)
	m["key"] = 10
	ch := make(chan int, 10)
	close(ch)

	var c chan int
	c = make(chan int, 10)
	close(c)
}

func arrayAndSlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	slice := arr[1:3]
	fmt.Println(slice)
	slice = append(slice, 6)
	fmt.Println(slice)
	slice = append(slice, 7)
	fmt.Println(slice)
	fmt.Sprintf("slice: %p\n", slice)
}

// 值传递，数组是值类型，修改不会影响原数组
func modifyArray(arr [3]int) {
	arr[0] = 10
}

// 引用传递，切片是引用类型，修改会影响原切片
func modifySlice(slice []int) {
	slice[0] = 10
}

func main() {
}

// abcabcbb
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	set := make(map[byte]bool, len(s))
	l, r, ans := 0, 0, 0
	for r < len(s) {
		//如果当前窗口里有重复字符
		if _, ok := set[s[r]]; ok {
			//移动左指针直到没有重复字符；更新答案
			for l < r {
				delete(set, s[l])
				l++
				ans = max(ans, r-l)

				//再判断一下当前还是否重复，不重复就直接break
				if _, ok := set[s[r]]; !ok {
					break
				}
			}
		}
		set[s[r]] = true
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var virtualHead *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = virtualHead
		virtualHead = cur
		cur = next
	}
	return virtualHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
