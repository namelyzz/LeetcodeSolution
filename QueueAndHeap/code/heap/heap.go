package heap

import "fmt"

type MaxHeap struct {
	slice []int
}

func (h *MaxHeap) Push(val int) {
	h.slice = append(h.slice, val)
	h.siftUp(len(h.slice) - 1)
}

func (h *MaxHeap) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.slice[index] <= h.slice[parent] {
			break
		}
		h.slice[index], h.slice[parent] = h.slice[parent], h.slice[index]
		index = parent
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.slice) == 0 {
		fmt.Println("heap is empty")
		return -1
	}

	top := h.slice[0]
	h.slice[0] = h.slice[len(h.slice)-1]
	h.slice = h.slice[:len(h.slice)-1]
	h.siftDown(0)
	return top
}

func (h *MaxHeap) siftDown(index int) {
	size := len(h.slice)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < size && h.slice[left] > h.slice[largest] {
			largest = left
		}
		if right < size && h.slice[right] > h.slice[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.slice[index], h.slice[largest] = h.slice[largest], h.slice[index]
		index = largest
	}
}

func (h *MaxHeap) PrintHeap() {
	for i := 0; i < len(h.slice); i++ {
		fmt.Print(h.slice[i], " ")
	}
	fmt.Println()
}

func Heapify(arr []int) *MaxHeap {
	h := &MaxHeap{slice: arr}
	for i := len(arr) / 2; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}
