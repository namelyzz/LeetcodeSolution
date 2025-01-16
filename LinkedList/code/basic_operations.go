package linklist

import "fmt"

/*
插入节点
*/

// insertAtHead 在头部插入节点
func insertAtHead(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val, Next: head}
	return newNode
}

// insertAtTail 在尾部插入节点
func insertAtTail(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val, Next: head}
	if head == nil {
		return newNode
	}

	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = newNode
	return head
}

// insertAfterKthNode 在指定位置（第 k 个节点后）插入节点
func insertAfterKthNode(head *ListNode, val int, k int) *ListNode {
	if k < 0 {
		return head
	}

	newNode := &ListNode{Val: val, Next: head}
	if k == 0 {
		newNode.Next = head
		return newNode
	}

	cur := head
	for i := 0; i < k-1 && cur != nil; i++ {
		cur = cur.Next
	}
	if cur != nil {
		newNode.Next = cur.Next
		cur.Next = newNode
	}

	return head
}

/*
删除节点
*/

// 删除头节点
func deleteHead(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return head.Next
}

// 删除尾节点
func deleteTail(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	cur := head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head
}

// 删除指定值的节点
func deleteNodeWithValue(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	cur := head

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return dummy.Next
}

// 删除指定位置的节点
func deleteNodeAtPosition(head *ListNode, k int) *ListNode {
	if k < 0 || head == nil {
		return head
	}

	if k == 0 {
		return head.Next
	}

	cur := head
	for i := 0; i < k-1 && cur != nil; i++ {
		cur = cur.Next
	}
	if cur != nil && cur.Next != nil {
		cur.Next = cur.Next.Next
	}

	return head
}

/*
查找节点
*/

// 按值查找节点
func findNodeByValue(head *ListNode, val int) *ListNode {
	cur := head
	for cur != nil {
		if cur.Val == val {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

// 按索引遍历链表
func traverseByIndex(head *ListNode, index int) *ListNode {
	if index < 0 {
		return nil
	}

	cur := head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	return cur
}

func BasicOperate() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}

	head = insertAtHead(head, 0)
	fmt.Println("头部插入节点")
	printList(head)

	head = insertAtTail(head, 4)
	fmt.Println("尾部插入节点")
	printList(head)

	head = insertAfterKthNode(head, 5, 2)
	fmt.Println("在指定位置插入节点")
	printList(head)

	// 删除节点
	//head = deleteHead(head)
	//printList(head)
	//
	//head = deleteTail(head)
	//printList(head)
	//
	//head = deleteNodeWithValue(head, 2)
	//printList(head)
	//
	//head = deleteNodeAtPosition(head, 1)
	//printList(head)
	//
	//// 查找节点
	//node := findNodeByValue(head, 3)
	//if node != nil {
	//	fmt.Println("Found node with value 3:", node.Val)
	//} else {
	//	fmt.Println("Node not found")
	//}
	//
	//node = traverseByIndex(head, 2)
	//if node != nil {
	//	fmt.Println("Node at index 2:", node.Val)
	//} else {
	//	fmt.Println("Index out of range")
	//}
}
