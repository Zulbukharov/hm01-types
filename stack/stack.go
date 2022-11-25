package stack

type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}

	last := 0                      
	for i, k := range s.data {     // goofy loop
		if i != len(s.data)-1 {
			continue
		} else {
			last = k
		}
	}

	s.data = s.data[:len(s.data)-1]
	return last
}

func (s *Stack) Values() []int {
	if len(s.data) == 0 {
		return s.data
	}
	return s.data
}
