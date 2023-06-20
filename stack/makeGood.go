package stack

func MakeGood(s string) string {
	if len(s) < 2 {
		return s
	}
	stack := new(Stack)
	stack.Push(s[0])
	for i := 1; i < len(s); i++ {
		c := s[i]
		t, empty := stack.Pop()
		if empty {
			stack.Push(c)
			continue
		}
		if c-32 == t || t-32 == c {
			continue
		}
		stack.Push(t)
		stack.Push(c)
	}
	return string(stack.Items)
}
