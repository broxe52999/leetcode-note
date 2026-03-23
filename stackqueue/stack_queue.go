package stackqueue

import (
	"errors"
	"leetcode/stack"
)

// Queue 使用两个栈实现的队列
type Queue struct {
	inStack  stack.Stack // 入队栈：只负责接收新元素
	outStack stack.Stack // 出队栈：只负责弹出元素
}

func NewQueue() *Queue {
	return &Queue{
		inStack:  *stack.NewStack(),
		outStack: *stack.NewStack(),
	}
}

// Enqueue 入队：直接压入 inStack，O(1)
func (q *Queue) Enqueue(val interface{}) {
	q.inStack.Push(val)
}

// Dequeue 出队：从 outStack 弹出，若 outStack 空则将 inStack 全部倒入
func (q *Queue) Dequeue() (interface{}, error) {
	// 如果 outStack 为空，将 inStack 全部倒入
	if q.outStack.IsEmpty() {
		for !q.inStack.IsEmpty() {
			val, _ := q.inStack.Pop()
			q.outStack.Push(val)
		}
	}

	// 如果倒完后 outStack 仍为空，说明队列整体为空
	if q.outStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	return q.outStack.Pop()
}

// Peek 查看队首元素
func (q *Queue) Peek() (interface{}, error) {
	if q.outStack.IsEmpty() {
		for !q.inStack.IsEmpty() {
			val, _ := q.inStack.Pop()
			q.outStack.Push(val)
		}
	}
	if q.outStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.outStack.Top()
}

func (q *Queue) IsEmpty() bool {
	return q.inStack.IsEmpty() && q.outStack.IsEmpty()
}
