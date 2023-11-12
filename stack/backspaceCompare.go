package stack

func BackspaceCompare(s string, t string) bool {
	var stack1, stack2 Stack

	for _, n := range s {
		if n != '#' {
			stack1.Push(byte(n))
		} else {
			stack1.Pop()
		}
	}

	for _, n := range t {
		if n != '#' {
			stack2.Push(byte(n))
		} else {
			stack2.Pop()
		}
	}

	return string(stack1.Items) == string(stack2.Items)
}
