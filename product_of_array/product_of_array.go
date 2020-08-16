package product_of_array

// Use only two loops for performance: O(n)
// strategy: multiply all and divide by n
func ProductExceptSelf(nums []int) []int {
	newNums := make([]int, len(nums))
	indexOfZeroes := []int{}

	productOfAll := 1
	for i, n := range nums {
		if n == 0 {
			// Collect indexes of all zeroes
			indexOfZeroes = append(indexOfZeroes, i)
		} else {
			productOfAll *= n
		}
	}

	for i, n := range nums {
		// if there are more than one zero, then the product result will always be zero
		if len(indexOfZeroes) > 1 ||
			// if there is one zero, and the index is not on that, then the product will be zero
			(len(indexOfZeroes) == 1 && indexOfZeroes[0] != i) {
			newNums[i] = 0
			continue
		}
		// in case the n is zero and divider cannot be zero
		if n != 0 {
			nn := productOfAll / n
			newNums[i] = nn
		} else {
			newNums[i] = productOfAll
		}
	}
	return newNums
}
