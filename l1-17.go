package main

func binarySearch(nums []int, el int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == el {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == el {
			return middle
		}

		if nums[middle] > el {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}