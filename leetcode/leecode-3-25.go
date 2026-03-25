package leetcode

import (
	"sort"
)

// 1. 三数之和
// 找到所有三个数和为0的组合，不能重复使用某元素
// 暴力，N3
//  N2
//
/*
如果判断第三个数的时候，用哈希表，可能出现一个问题
比如  1 -2 0 这个数组，其实无解，但是用哈希表，
i = 1  j = -2 时候， 找k  发现1确实存在，然后给出一个答案 1 -2 1这不就错了吗？

题解：
要点： 排序+ 双指针
先排序，[0, 1, 2, 2, 2, 3]
每个元素都是递增的，然后比如枚举到2的时候，要跳到3
就能保证不重复
如果我们固定了前两重循环枚举到的元素 a 和 b，
那么只有唯一的 c 满足 a+b+c=0。
当第二重循环往后枚举一个元素 b2时，由于 b2>b，
那么满足 a+b2 + c2 =0 的c2一定有 c2<c，
即 c2在数组中一定出现在 c 的左侧。
也就是说，我们可以从小到大枚举 b，
同时从大到小枚举 c，
即第二重循环和第三重循环实际上是并列的关系。

有了这样的发现，我们就可以保持第二重循环不变，
而将第三重循环变成一个从数组最右端开始向左移动的指针，

关键点在于
k 的指针不是在j的循环里，而是在j的循环外面，是并列的
在a,b 固定下，c肯定是唯一的,假设找到了一个c 使得 a+b+c = 0
那么下一个答案，一定是b开始增大，因为内层循环先动，b往右增大，
c呢 如果还在c的位置或者c往右的位置，肯定不满足a+b+c=0
所以c的位置，从始至终，只能往左移动，直到j == k
*/

/*
数组排序
遍历每个元素作为第一个数 a
如果 a > 0：后面都是更大的数，不可能和为 0，直接退出
如果当前数和前一个相同，跳过（去重）
左指针 left = i+1，右指针 right = len(nums)-1
计算三数之和：
和 = 0：记录结果，同时移动左右指针并跳过重复值
和 < 0：左指针右移，让和变大
和 > 0：右指针左移，让和变小
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var left, right int
		left = i + 1
		right = n - 1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			// == target left++ right --,
			if sum == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++

			} else {
				right--
			}

		}
	}
	return ans
}

/*
2.   最大子数组和
给你一个整数数组 nums ，
请你找出一个具有最大和的连续子数组
（子数组最少包含一个元素），返回其最大和。


关键：dp[i] 表示以i结尾的最大子数组和
dp[i] = (nums[i] + dp[i-1], nums[i])
ans:= max(dp[i])
*/

func maxSubArray(nums []int) int {
	n := len(nums)
	ans := nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		ans = max(ans, sum)
	}
	return ans
}

/*
3. 最长回文子串

给你一个字符串 s，找到 s 中最长的 回文 子串。
输入：s = "babad"
输出："bab"

最优通用解法：中心扩散法
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	n := len(s)
	maxLen := 0
	var ans string
	for i := 0; i < n; i++ {
		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		len1 := r1 - l1 + 1
		len2 := r2 - l2 + 1
		if len1 > maxLen {
			maxLen = len1
			ans = s[l1 : r1+1]
		}
		if len2 > maxLen {
			maxLen = len2
			ans = s[l2 : r2+1]
		}

	}
	return ans
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	//退出循环多加了一次
	return left + 1, right - 1
}

// 4. 合并两个有序链表
/*
两个链表本身已经有序，所以不用排序，只需要一条链串起来就行：
用一个虚拟头节点（dummy） 方便拼接
用一个指针 cur 从头往后走
每次比较两个链表当前节点的值
谁小就把谁接在 cur 后面
然后对应链表指针后移
最后把剩下没走完的链表直接接在尾部

*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next

		}
		cur = cur.Next

	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

/*
5. 二叉树的层序遍历

给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
（即逐层地，从左到右访问所有节点）。

我觉得核心就是记录每层的节点个数
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		levelNode := []int{}
		levelSize := len(queue) // 当前层有多少节点

		for i := 0; i < levelSize; i++ {

			node := queue[0]
			queue = queue[1:]

			levelNode = append(levelNode, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelNode)
	}
	return res
}

/*
6.
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，
请你计算网格中岛屿的数量。

岛屿总是被水包围，
并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

题解：
走过的地方标记为-1；
一开始先找到1的位置，进去，进行dfs
*/
func numIslands(grid [][]byte) int {
	ans := 0
	m := len(grid)
	n := len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i-1 >= 0 && grid[i-1][j] == '1' {
			grid[i-1][j] = '0'
			dfs(i-1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			grid[i][j-1] = '0'
			dfs(i, j-1)
		}
		if j+1 <= n-1 && grid[i][j+1] == '1' {
			grid[i][j+1] = '0'
			dfs(i, j+1)
		}
		if i+1 <= m-1 && grid[i+1][j] == '1' {
			grid[i+1][j] = '0'
			dfs(i+1, j)
		}
	}
	//m*n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans

}

/*
7. 买卖股票的最佳时机

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

核心：
记录一个此前的历史最低价格，每天卖出的利润就是当天减去历史最低价格，
然后更新最大利润和最低价格
*/
func maxProfit(prices []int) int {
	ans := 0
	if len(prices) == 1 {
		return ans
	}
	lowest := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-lowest > ans {
			ans = prices[i] - lowest
		}
		if prices[i] < lowest {
			lowest = prices[i]
		}
	}
	return ans

}

/*
8.二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，\
最近公共祖先表示为一节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大
（一个节点也可以是它自己的祖先）。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归终止：空节点 或 找到p/q
	if root == nil || root == q || root == p {
		return root
	}
	// 左右分别查找
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左边空 → 答案在右边
	if left == nil {
		return right
	}
	// 右边空 → 答案在左边
	if right == nil {
		return left
	}
	// 都不空 → 当前节点就是LCA
	return root
}

/*
9. 反转链表 II
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	// 1. 走到 left 前一个节点,找到前驱节点
	//走left-1步
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// cur 待反转区间的第一个节点
	cur := pre.Next
	// 2. 反转区间：一共要反转 right-left 次
	for i := 0; i < right-left; i++ {
		// 要拎到前面的节点
		next := cur.Next
		// 把 next 从原位置摘走
		cur.Next = next.Next
		// 插到 pre 后面
		next.Next = pre.Next
		// 更新 pre 的下一个
		pre.Next = next

	}
	return dummy.Next
}

/*
10. 全排列
给定一个不含重复数字的数组 nums ，
返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

//dfs
*/
func permute(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	used := make([]bool, n)
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])

			dfs(path)

			path = path[:len(path)-1]
			used[i] = false

		}
	}

	dfs([]int{})
	return res
}

/*
11. 搜索旋转排序数组

整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）
上进行了 向左旋转，使数组变为
[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，
如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 左边有序
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

/*
12.二叉树的锯齿形层序遍历
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。
（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	flag := true //true向左，false向右

	for len(queue) != 0 {
		level := []int{}
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if !flag {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		flag = !flag
		res = append(res, level)
	}
	return res
}

/*
13. 二叉树的右视图
*/
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

/*
14.最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，
删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

我的想法：
dp[i] 表示以i 结尾的 最长递增子序列长度
- 对每个 i，遍历前面所有 j < i：
  - if nums[i] > nums[j]，则 dp[i] = max(dp[i], dp[j]+1)
*/
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, num := range nums {
		// 二分找第一个 >= num 的下标
		l, r := 0, len(tails)
		for l < r {
			mid := (l + r) / 2
			if tails[mid] >= num {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l == len(tails) {
			tails = append(tails, num)
		} else {
			tails[l] = num
		}
	}
	return len(tails)
}

/*
15.螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix
，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

我的思路，用dfs
// 通过右下左上的顺序一致访问到边界，访问过的标记为访问
*/
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ans := make([]int, 0)
	i, j := 0, 0
	ans = append(ans, matrix[i][j])
	matrix[i][j] = 101
	for len(ans) != m*n {
		//右
		for j+1 < n && matrix[i][j+1] != 101 {
			ans = append(ans, matrix[i][j+1])
			matrix[i][j+1] = 101
			j++

		}
		//下
		for i+1 < m && matrix[i+1][j] != 101 {
			ans = append(ans, matrix[i+1][j])
			matrix[i+1][j] = 101
			i++

		}
		for j >= 1 && matrix[i][j-1] != 101 {
			ans = append(ans, matrix[i][j-1])
			matrix[i][j-1] = 101
			j--

		}
		for i >= 1 && matrix[i-1][j] != 101 {
			ans = append(ans, matrix[i-1][j])
			matrix[i-1][j] = 101
			i--

		}
	}
	return ans
}
