package stack

type Stack struct {
	values []int
}

func (s *Stack) Push(v int) {
	s.values = append([]int{v}, s.values...)

}
func (s *Stack) Pop() int {
	if len(s.values) == 0 {
		return 0
	}
	v := s.values[0]
	s.values = s.values[1:]
	return v
}
func (s *Stack) Values() []int {

	if len(s.values) == 0 {
		return nil
	}
	return s.values
}
