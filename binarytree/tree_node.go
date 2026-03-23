package binarytree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// 前序遍历： 根左右
// 递归版本
func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 1. 访问根节点
	fmt.Printf("%d ", root.Val)
	// 2. 递归遍历左子树
	PreorderTraversal(root.Left)
	// 3. 递归遍历右子树
	PreorderTraversal(root.Right)
}

// 迭代版本（使用栈）
// 切片模拟栈
// PreorderTraversalIterative 迭代前序遍历
func PreorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)

	// 用切片模拟栈
	stack := make([]*TreeNode, 0)

	// 初始化：将根节点压栈
	stack = append(stack, root)

	//栈不为空的时候
	for len(stack) > 0 {
		// 弹出栈顶元素
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问当前节点
		result = append(result, node.Val)

		// 注意：因为栈是后进先出，要实现根左右，需要先压右孩子，再压左孩子
		// 这样出栈时才是左孩子先出
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 中序遍历： 左根右
func InorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 1. 递归遍历左子树
	InorderTraversal(root.Left)
	// 2. 访问根节点
	fmt.Printf("%d ", root.Val)
	// 3. 递归遍历右子树
	InorderTraversal(root.Right)
}

// InorderTraversalIterative 迭代中序遍历
func InorderTraversalIterative(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		// 1. 一直向左走，将沿途节点压栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 2. 弹出栈顶节点（最左边的节点）
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 3. 访问该节点
		result = append(result, current.Val)

		// 4. 转向右子树
		current = current.Right
	}

	return result
}

// PostorderTraversal 递归后序遍历
func PostorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 1. 递归遍历左子树
	PostorderTraversal(root.Left)
	// 2. 递归遍历右子树
	PostorderTraversal(root.Right)
	// 3. 访问根节点
	fmt.Printf("%d ", root.Val)
}

// PostorderTraversalIterative 迭代后序遍历（单栈+标记法）
func PostorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisited *TreeNode // 记录上一次访问的节点

	current := root

	for current != nil || len(stack) > 0 {
		// 1. 一直向左走，将沿途节点压栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 查看栈顶节点
		peek := stack[len(stack)-1]

		// 2. 如果右子树存在且未被访问过，先处理右子树
		if peek.Right != nil && lastVisited != peek.Right {
			current = peek.Right
		} else {
			// 3. 否则，访问当前节点
			result = append(result, peek.Val)
			lastVisited = peek
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

// PostorderTraversalTwoStack 迭代后序遍历（双栈法，更易理解）
func PostorderTraversalTwoStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	s1 := make([]*TreeNode, 0) // 辅助栈
	s2 := make([]*TreeNode, 0) // 输出栈

	s1 = append(s1, root)

	for len(s1) > 0 {
		// 弹出s1栈顶，压入s2
		node := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, node)

		// 先左后右入s1（这样出s1时是先右后左，入s2后变成根右左，最后反转成左右根）
		if node.Left != nil {
			s1 = append(s1, node.Left)
		}
		if node.Right != nil {
			s1 = append(s1, node.Right)
		}
	}

	// 弹出s2即为后序遍历结果
	for i := len(s2) - 1; i >= 0; i-- {
		result = append(result, s2[i].Val)
	}

	return result
}
