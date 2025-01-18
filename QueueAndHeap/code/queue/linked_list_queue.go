package queue

import (
	"errors"
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedListQueue struct {
	front *ListNode // 队首节点
	rear  *ListNode // 队尾节点
}

func NewLinkedListQueue() *LinkedListQueue {
	dummy := &ListNode{}
	return &LinkedListQueue{
		front: dummy,
		rear:  dummy,
	}
}

func (q *LinkedListQueue) Enqueue(val int) {
	newNode := &ListNode{val: val}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *LinkedListQueue) Dequeue() (int, error) {
	if q.front.next == nil {
		return -1, errors.New("queue is empty")
	}

	val := q.front.next.val
	q.front.next = q.front.next.next
	if q.front.next == nil {
		q.rear = q.front // 队列清空时重置rear
	}
	return val, nil
}

func (q *LinkedListQueue) PrintQueue() {
	if q.front.next == nil {
		fmt.Println("队列为空")
		return
	}
	fmt.Print("队列中的元素为: ")
	current := q.front.next
	for current != nil {
		fmt.Print(current.val, " ")
		current = current.next
	}
	fmt.Println()
}
