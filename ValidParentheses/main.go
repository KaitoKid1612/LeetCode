package main

func isValid(s string) bool {
	stack := []rune{}

	// Create a map to store the matching pairs of parentheses
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if len(stack) > 0 && stack[len(stack)-1] == pairs[char] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	// Example usage
	s := "([{}])"
	result := isValid(s)
	println(result) // Output: true
}
