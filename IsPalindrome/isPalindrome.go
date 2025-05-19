package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false // Negative numbers are not palindromes
	}

	original := x
	reversed := 0

	for x > 0 {
		remainder := x % 10
		reversed = reversed*10 + remainder
		x /= 10
	}

	if original == reversed {
		return true
	}
	return false
}

func main() {
	// Example usage of isPalindrome
	result := isPalindrome(121)
	println(result) // Output: true
}
