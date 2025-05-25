package main

func longestPalindrome(words []string) int {
	palindromeCount := 0
	seen := make(map[string]int)

	for _, word := range words {
		runes := []rune(word)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		reverseWord := string(runes)
		if count, exists := seen[reverseWord]; exists && count > 0 {
			palindromeCount += 4
			seen[reverseWord]--
		} else {
			seen[word]++
		}
	}

	for word, count := range seen {
		if count > 0 && word[0] == word[1] {
			palindromeCount += 2
			break
		}
	}

	return palindromeCount
}

func main() {
	words := []string{"lc", "cl", "aa"}
	result := longestPalindrome(words)
	println(result)
}
