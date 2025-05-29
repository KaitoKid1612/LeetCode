package main

func strStr(haystack string, needle string) int {
	if needle == "" || needle == haystack {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	// Example usage
	haystack := "hello"
	needle := "ll"
	result := strStr(haystack, needle)
	println(result) // Output: 2
}
