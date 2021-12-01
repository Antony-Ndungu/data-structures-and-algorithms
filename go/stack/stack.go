package stack

type Stack struct {
	elements []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Push(data int) {
	s.elements = append(s.elements, data)
}

func (s *Stack) Pop() int {
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}
