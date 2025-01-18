package queue

import (
	"errors"
	"fmt"
)

type ArrayQueue struct {
	elements []int
	front    int // 队首索引
	rear     int // 队尾索引
	size     int // 队列中元素的个数
	capacity int // 队列的容量
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		elements: make([]int, capacity),
		capacity: capacity,
	}
}

func (q *ArrayQueue) Enqueue(val int) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}

	q.elements[q.rear] = val

	// 使用取模运算符 % 来更新 q.rear 的值。
	// 这是为了实现循环队列的功能，当 q.rear 达到队列的最大容量时，会重新从队列的头部开始。
	// 例如，如果队列容量为 5，当 q.rear 为 4 时，(4 + 1) % 5 的结果为 0，这样就可以继续在队列的头部插入元素。
	q.rear = (q.rear + 1) % q.capacity
	q.size++
	return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}

	val := q.elements[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return val, nil
}

func (q *ArrayQueue) PrintQueue() {
	if q.size == 0 {
		fmt.Println("队列为空")
		return
	}
	fmt.Print("队列中的元素为：")
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		fmt.Print(q.elements[index], " ")
	}
	fmt.Println()
}
