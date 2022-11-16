package stack

type Stack struct{}

func (s *Stack) Push(v int)    {}
func (s *Stack) Pop() int      { return 0 }
func (s *Stack) Values() []int { return nil }
