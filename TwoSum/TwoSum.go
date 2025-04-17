package main

func twoSum(nums []int, target int) []int {
	// Create a map to store the indices of the numbers
	numMap := make(map[int]int)
	// Iterate through the numbers
	for i, num := range nums {
		// Calculate the complement
		complement := target - num
		// Check if the complement exists in the map
		if index, found := numMap[complement]; found {
			// If found, return the indices
			return []int{index, i}
		}
		// Store the index of the current number in the map
		numMap[num] = i
	}
	// If no solution is found, return an empty slice
	return []int{}
	// Note: The problem guarantees that there is exactly one solution,
	// so we don't need to handle the case where no solution exists.
	// However, if needed, we could return an error or a specific value.
	// For example, we could return nil or a slice with -1 values.
	// return nil
}


func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)

	// Print the result
	for i, v := range result {
		println("Index", i, ":", v)
	}
}