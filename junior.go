package main

func isValid(s string) bool {
	stack := []string{}
	pairs := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, ch := range s {
		c := string(ch)
		switch c {
		case "(", "[", "{":
			stack = append(stack, c)
		case ")", "]", "}":
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != pairs[c] {
				return false
			}
		}
	}

	return len(stack) == 0
}
