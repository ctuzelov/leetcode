package stack

type Stack struct {
	Items []byte
}

func (s *Stack) Push(item byte) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() (byte, bool) {
	l := len(s.Items) - 1
	if l < 0 {
		return ' ', false
	}
	ToRemove := s.Items[l]
	s.Items = s.Items[:l]
	return ToRemove, true
}

func IsValid(s string) bool {
	var prn Stack
	for i := 0; i < len(s); i++ {
		t := s[i]
		if t == '{' || t == '[' || t == '(' {
			prn.Push(t)
			continue
		}
		valid, fl := prn.Pop()
		if !fl {
			return false
		}
		if t == ')' && valid != '(' {
			return false
		}
		if t == '}' && valid != '{' {
			return false
		}
		if t == ']' && valid != '[' {
			return false
		}
	}
	return len(prn.Items) == 0
}
