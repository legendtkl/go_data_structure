package ds

type Stack struct {
	value []interface{}
}

func (s *Stack) Empty() bool {
	return len(s.value) == 0
}
func (s *Stack) Top() interface{} {
	return s.value[len(s.value)-1]
}

func (s *Stack) Push(item interface{}) {
	s.value = append(s.value, item)
}

func (s *Stack) Pop() {
	s.value = s.value[:len(s.value)-1]
}
