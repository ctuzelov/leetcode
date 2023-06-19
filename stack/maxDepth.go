package stack

func (s *Stack) Len() int {
	return len(s.Items)
}

func MaxDepth(s string) int {
	var prn Stack
	max := 0
	for i := 0; i < len(s); i++ {
		t := s[i]
		if t == '(' {
			prn.Push(t)
			continue
		}
		if !(t == ')') {
			continue
		}
		l := prn.Len()
		if l > max {
			max = l
		}
		prn.Pop()
	}
	return max
}
