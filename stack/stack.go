package stack

import (
	"errors"
)

// Stack 定义栈结构，使用切片存储
type Stack struct {
	elements []interface{} // 使用空接口实现任意类型存储
}

// NewStack 初始化栈
func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

// Push 压入元素
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop 弹出栈顶元素
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	// 获取栈顶元素
	index := len(s.elements) - 1
	value := s.elements[index]
	// 弹出元素（通过切片截断）
	s.elements = s.elements[:index]
	return value, nil
}

// Top 获取栈顶元素但不弹出
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("栈为空")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size 获取栈大小
func (s *Stack) Size() int {
	return len(s.elements)
}
