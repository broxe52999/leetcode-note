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

// 双向链表，头插法，
// 双向链表的head 和tail 都是虚拟节点
// tail 节点的前一个结点代表最久未访问的元素
// get 直接获取
// put 若存在，则更新，对于双向链表，则是移动到head， 也就是先删除该节点，再在head添加新节点
// 若不存在，则直接在头部添加新节点， 若超出容量，则删除末尾节点
// 因此有几种操作链表的方法： removeNode 、 removeTail、 addHead、moveToHead
type DLinkedNode struct {
	key, value int
	pre, next  *DLinkedNode
}

func innitDlinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	cache      map[int]*DLinkedNode
	size       int
	capacity   int
	head, tail *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cache:    make(map[int]*DLinkedNode),
		capacity: capacity,
		size:     0,
		head:     innitDlinkedNode(0, 0),
		tail:     innitDlinkedNode(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.next = cache.head
	return cache
}

// 删除双向链表中间的一个元素
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

// 返回移除的节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node

}

// 将一个node 添加到头节点
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := innitDlinkedNode(key, value)
		this.addToHead(node)
		this.cache[key] = node
		this.size++
		if this.size > this.capacity {
			removedNode := this.removeTail()
			delete(this.cache, removedNode.key)
			this.size--
		}
	}
}

func findKthLargest(nums []int, k int) int {

	return 0
}

// 时间复杂度小于 O(n2)
// 思路：
// 1. 使用集合来存储每个数是否存在
// 2. 遍历nums，对于nums[i] ，只要看taget-nums[i] 是否存在即可
func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if j, ok := set[target-nums[i]]; ok {
			return []int{i, j}
		}
		set[nums[i]] = i
	}
	return nil
}

// 88. 合并两个有序数组
// 思路：
// 1.双指针，从尾部开始合并
// 2.直接合并的话，在nums1上从头开始合并会覆盖掉nums1的元素，所以考虑从尾部合并
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		for i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		}
		for i >= 0 && j >= 0 && nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// 20. 有效的括号
// 思路：
// 1. 遍历字符串，若遇到左括号就入栈
// 2. 遇到右括号，就和栈顶配对，若成功就弹出栈；若失败，则不是有效的括号字符串
// 3. 特殊边界： 最后一定要栈为空才为有效括号，不为空说明有多余的右括号
func isValid(s string) bool {
	runes := make([]rune, len(s)+1)
	tail := -1
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			tail++
			runes[tail] = v
		} else {
			if tail == -1 {
				return false
			}
			switch v {
			case ')':
				if runes[tail] == '(' {
					tail--
				} else {
					return false
				}
			case ']':
				if runes[tail] == '[' {
					tail--
				} else {
					return false
				}
			case '}':
				if runes[tail] == '{' {
					tail--
				} else {
					return false
				}
			}
		}
	}
	return tail == -1
}

func main() {
	// 创建一个切片作为栈
	stack := make([]int, 0, 5)
	stack = append(stack, 1, 2, 3, 4, 5)

	fmt.Printf("初始状态:\n")
	fmt.Printf("stack = %v\n", stack)
	fmt.Printf("len = %d, cap = %d\n\n", len(stack), cap(stack))

	// 查看底层数组（通过反射，仅用于演示）
	fmt.Printf("底层数组中的实际值: %v\n\n", stack[:cap(stack)]) // 访问整个底层数组

	// 执行 Pop 操作
	fmt.Printf("执行: stack = stack[:len(stack)-1]\n")
	stack = stack[:len(stack)-1]

	fmt.Printf("Pop 后:\n")
	fmt.Printf("stack = %v\n", stack)
	fmt.Printf("len = %d, cap = %d\n\n", len(stack), cap(stack))

	fmt.Printf("底层数组现在: %v\n", stack[:cap(stack)])
	fmt.Printf("注意: 元素5仍然在底层数组中！\n\n")

	// 再 Push 一个元素
	fmt.Printf("执行 Push 6:\n")
	stack = append(stack, 6)

	fmt.Printf("Push 后:\n")
	fmt.Printf("stack = %v\n", stack)
	fmt.Printf("底层数组现在: %v\n", stack[:cap(stack)])
	fmt.Printf("元素5被覆盖了！\n")
}
