package slidingwindow

//1. Find the Maximum Sum of a Subarray of Size K
//
//Problem:
//Given an array of integers and an integer K, find the maximum sum of any contiguous subarray of size K.
//
//Example:
//
//Input: arr = [2, 1, 5, 1, 3, 2], K = 3
//Output: 9
//Explanation: The subarrays of size 3 are:
//[2, 1, 5] → sum = 8
//[1, 5, 1] → sum = 7
//[5, 1, 3] → sum = 9 (maximum)
//[1, 3, 2] → sum = 6
//
//Sliding Window Approach:
//
//Start with the sum of the first K elements.
//Slide the window by removing the first element and adding the next one.
//Keep track of the maximum sum encountered.

func MaxSumOfSubArraySizeK(nums []int, k int) int {
	n := len(nums)
	// edge case: array smaller than k
	if k > n {
		return -1
	}

	maxSum, windowSum := 0, 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum = windowSum

	for start, end := 0, k; end < n; start, end = start+1, end+1 {
		windowSum += nums[end] - nums[start] // Add new element, remove old
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}
