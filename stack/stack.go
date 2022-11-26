package stack

type Stack struct {
	Obj []int
}

func (s *Stack) Push(v int) {
	s.Obj = append(s.Obj, v)
}

func (s *Stack) Pop() int {
	if len(s.Obj) == 0 {
		return 0
	}
	v := s.Obj[len(s.Obj)-1]
	s.Obj = s.Obj[:len(s.Obj)-1]
	return v
}

func (s *Stack) Values() []int {
	if len(s.Obj) == 0 {
		return s.Obj
	}
	return s.Obj
}
