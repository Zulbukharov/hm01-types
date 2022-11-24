package stack

type Stack struct {
	values []int
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)

}
func (s *Stack) Pop() int {
	if len(s.values) == 0 {
		return 0
	}
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}
func (s *Stack) Values() []int {

	if len(s.values) == 0 {
		return s.values
	}
	return s.values
}
