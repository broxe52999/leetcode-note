package stack

import "errors"

// Node 链表节点
type Node struct {
	value interface{}
	next  *Node
}

// LinkedStack 链式栈
type LinkedStack struct {
	top  *Node // 栈顶指针
	size int   // 栈大小
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{
		top:  nil,
		size: 0,
	}
}

// Push 入栈（头插法）
func (s *LinkedStack) Push(value interface{}) {
	newNode := &Node{
		value: value,
		next:  s.top, // 新节点指向旧的栈顶
	}
	s.top = newNode // 更新栈顶为新节点
	s.size++
}

// Pop 出栈
func (s *LinkedStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	value := s.top.value
	s.top = s.top.next // 栈顶移动到下一个节点
	s.size--
	return value, nil
}

// Pop 出栈
func (s *LinkedStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	value := s.top.value
	return value, nil
}

// IsEmpty ...
func (s *LinkedStack) IsEmpty() bool {
	return s.top == nil
}
