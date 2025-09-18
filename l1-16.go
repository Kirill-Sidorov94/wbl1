package main

import(
	"math/rand"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]

	left := 0
	right := len(nums) - 1

	for i := 0; i <= right; {
	    if nums[i] < pivot {
	        nums[left], nums[i] = nums[i], nums[left]
	        left++
	        i++
	    } else if nums[i] > pivot {
	        nums[right], nums[i] = nums[i], nums[right]
	        right--
	    } else {
	        i++
	    }
	}

	return append(quickSort(nums[:left]), append(nums[left:right+1], quickSort(nums[right+1:])...)...)
}