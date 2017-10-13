package datastructures

type stack struct {
	stack []int
	size int
}
func CreateStack() stack {
	stack := stack{ stack: make([]int, 10)}
	return stack
}
func (s *stack) Push(value int) int {
	s.stack[s.size] = value
	s.size++
	return value
}

func (s *stack) Pop() (int, bool)  {
	if s.size == 0 {
		return 0, false
	}
	s.size--
	val := s.stack[s.size]
	return val, true
}