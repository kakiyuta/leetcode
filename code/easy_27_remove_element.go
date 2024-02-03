/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

package code

// @lc code=start
func removeElement(nums []int, val int) int {
	var removiedNums []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			removiedNums = append(removiedNums, nums[i])
		}
	}

	for i := 0; i < len(removiedNums); i++ {
		nums[i] = removiedNums[i]
	}

	return len(removiedNums)
}

// @lc code=end
