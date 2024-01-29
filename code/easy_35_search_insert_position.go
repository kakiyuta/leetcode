/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package code

import "fmt"

// @lc code=start
func searchInsert(nums []int, target int) int {
	return searchInsertRecursive(nums, target, 0, len(nums))
}

func searchInsertRecursive(nums []int, target, start, end int) int {
	fmt.Println("debug")
	if start >= end {
		return start
	}

	mid := start + (end-start)/2
	fmt.Println("hoge")
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return searchInsertRecursive(nums, target, mid+1, end)
	} else {
		return searchInsertRecursive(nums, target, start, mid)
	}
}

// @lc code=end
