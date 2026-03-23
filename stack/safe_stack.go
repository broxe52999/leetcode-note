package stack

import (
	"errors"
	"sync"
)

// SafeStack 线程安全的栈
type SafeStack struct {
	elements []interface{}
	lock     sync.RWMutex // 使用读写锁
}

func NewSafeStack() *SafeStack {
	return &SafeStack{
		elements: make([]interface{}, 0),
		// lock: sync.RWMutex{},  // 零值可用，不需要显式初始化
	}
}
func (s *SafeStack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, value)
}

func (s *SafeStack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.elements) == 0 {
		return nil, errors.New("栈为空")
	}
	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val, nil
}

// IsEmpty 检查栈是否为空
func (s *SafeStack) IsEmpty() bool {
	if s == nil {
		return true
	}

	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.elements) == 0
}

// Top 查看栈顶元素但不弹出
func (s *SafeStack) Top() (interface{}, error) {
	if s == nil {
		return nil, errors.New("栈未初始化")
	}

	s.lock.RLock() // 读操作使用读锁
	defer s.lock.RUnlock()

	if len(s.elements) == 0 {
		return nil, errors.New("栈为空")
	}

	return s.elements[len(s.elements)-1], nil
}

// Size 获取栈大小
func (s *SafeStack) Size() int {
	if s == nil {
		return 0
	}

	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.elements)
}

// Clear 清空栈
func (s *SafeStack) Clear() {
	if s == nil {
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = make([]interface{}, 0) // 重新分配新的切片，让GC回收旧的
}
