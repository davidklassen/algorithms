package trees_and_graphs

import (
	"errors"
)

type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func shiftUp(arr []int) {
	val := arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; {
		p := (i - 1) / 2
		if arr[p] > val {
			tmp := arr[p]
			arr[p] = val
			arr[i] = tmp
		}
		i = p
	}
}

func shiftDown(arr []int) {
	min, i := 0, 0
	for {
		left, right := 2*i+1, 2*i+2
		if left < len(arr)-1 && arr[min] > arr[left] {
			min = left
		}
		if right < len(arr)-1 && arr[min] > arr[right] {
			min = right
		}
		if min != i {
			tmp := arr[min]
			arr[min] = arr[i]
			arr[i] = tmp
			i = min
		} else {
			break
		}
	}

}

func (h *Heap) Push(val int) {
	h.arr = append(h.arr, val)
	shiftUp(h.arr)
}

func (h *Heap) Pop() (res int, err error) {
	if h.IsEmpty() {
		return res, errors.New("cannot pop from empty heap")
	}
	res = h.arr[0]
	if len(h.arr) > 1 {
		h.arr[0] = h.arr[len(h.arr)-1]
	}
	h.arr = h.arr[:len(h.arr)-1]
	if !h.IsEmpty() {
		shiftDown(h.arr)
	}
	return res, nil
}

func (h *Heap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("cannot peek at empty heap")
	}
	return h.arr[0], nil
}

func (h *Heap) IsEmpty() bool {
	return len(h.arr) == 0
}
