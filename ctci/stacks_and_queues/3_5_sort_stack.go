package stacks_and_queues

import (
	"fmt"
)

func stackLen(s *Stack) int {
	res := 0
	tmp := NewStack()
	for !s.IsEmpty() {
		res++
		val, _ := s.Pop()
		tmp.Push(val)
	}
	for !tmp.IsEmpty() {
		val, _ := tmp.Pop()
		s.Push(val)
	}
	return res
}

func SortStack(s *Stack) {
	tmp := NewStack()
	sLen := stackLen(s)
	fmt.Println(sLen)
	for sLen != 0 {
		max, _ := s.Peek()
		maxCount := 0
		for i := 0; i < sLen; i++ {
			val, _ := s.Pop()
			if val > max {
				maxCount = 1
				max = val
			} else if val == max {
				maxCount++
			}
			tmp.Push(val)
		}
		for i := 0; i < maxCount; i++ {
			s.Push(max)
		}
		for !tmp.IsEmpty() {
			if val, _ := tmp.Pop(); val != max {
				s.Push(val)
			}
		}
		sLen -= maxCount
	}
}
