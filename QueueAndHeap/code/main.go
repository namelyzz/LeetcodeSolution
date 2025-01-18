package main

import (
    "fmt"
    "github.com/namelyzz/LeetcodeSolution/QueueAndHeap/code/heap"
    "github.com/namelyzz/LeetcodeSolution/QueueAndHeap/code/queue"
)

func main() {
    checkArrayQueue()
    fmt.Println("================================================")
    checkLinkedListQueue()
    fmt.Println("================================================")
    checkHeap()
}

func checkArrayQueue() {
    // 创建一个容量为 5 的队列
    arrQueue := queue.NewArrayQueue(5)

    // 入队操作
    fmt.Println("入队操作开始：")
    for i := 1; i <= 3; i++ {
        err := arrQueue.Enqueue(i)
        if err != nil {
            fmt.Printf("入队元素 %d 失败：%v\n", i, err)
        } else {
            fmt.Printf("入队元素 %d 成功\n", i)
            arrQueue.PrintQueue()
        }
    }

    // 出队操作
    fmt.Println("\n出队操作开始：")
    val, err := arrQueue.Dequeue()
    if err != nil {
        fmt.Printf("出队失败：%v\n", err)
    } else {
        fmt.Printf("出队元素 %d 成功\n", val)
        arrQueue.PrintQueue()
    }
}

func checkLinkedListQueue() {
    // 创建一个新的链表队列
    linkedQueue := queue.NewLinkedListQueue()

    // 入队操作
    fmt.Println("入队操作开始:")
    for i := 1; i <= 3; i++ {
        linkedQueue.Enqueue(i)
        fmt.Printf("入队元素 %d 后，", i)
        linkedQueue.PrintQueue()
    }

    // 出队操作
    fmt.Println("\n出队操作开始:")
    for i := 0; i < 4; i++ {
        val, err := linkedQueue.Dequeue()
        if err != nil {
            fmt.Printf("出队失败: %v\n", err)
        } else {
            fmt.Printf("出队元素 %d 后，", val)
            linkedQueue.PrintQueue()
        }
    }
}

func checkHeap() {
    h := heap.Heapify([]int{3, 1, 5, 4, 8, 9, 2})
    fmt.Println("当前堆元素：")
    h.PrintHeap()
    h.Push(7)
    fmt.Println("当前堆元素：")
    h.Push(0)
    fmt.Println("当前堆元素：")
    h.PrintHeap()
    fmt.Println("堆顶元素:", h.Pop())
    /*
       最终堆元素：9 8 5 7 1 3 2 4 0
       堆的样子：
              9
            /   \
           8     5
          / \   / \
         7   1 3   2
        / \
       4   0
    */
}
