package stack

type Stack struct {
	StackInt []int
}

func (s *Stack) Push(v int) {
	s.StackInt = append(s.StackInt, v)
}
func (s *Stack) Pop() int {
	if len(s.StackInt) == 1 {
		res := s.StackInt[0]
		s.StackInt = []int{}
		return res
	}
	if len(s.StackInt) > 1 {
		n := len(s.StackInt) - 1
		res := s.StackInt[n]
		s.StackInt = s.StackInt[:n]
		return res
	}
	return 0
}
func (s *Stack) Values() []int {
	return s.StackInt
}
