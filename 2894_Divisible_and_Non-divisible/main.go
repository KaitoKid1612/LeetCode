package main

func differrenOfSums(n int, m int) int {
	sumDivisible := 0
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			sumDivisible += i
		}
	}

	sumNonDivisible := 0
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			sumNonDivisible += i
		}
	}

	return sumNonDivisible - sumDivisible
}

func main() {
	// Example usage
	n := 10
	m := 3
	result := differrenOfSums(n, m)
	println(result) // Output: 19
}
