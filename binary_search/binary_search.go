package binary_search

func BinarySearch(nums []int, target int) int {
	index := -1

	left := 0
	right := len(nums) - 1

	for left <= right {
		midIndex := (left + right) / 2

		if nums[midIndex] == target {
			return midIndex
		}

		if nums[midIndex] < target {
			left = midIndex + 1
		}

		if nums[midIndex] > target {
			right = midIndex - 1
		}
	}

	return index
}
