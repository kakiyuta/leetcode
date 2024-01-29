/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package code

// @lc code=start
func searchInsert(nums []int, target int) int {
	return searchInsertRecursive(nums, target, 0)
}

// 再起的にインサートするindexを見つけ返す
func searchInsertRecursive(nums []int, target int, index int) int {
	// 配列の長さが1の場合
	if len(nums) == 1 {
		if nums[0] < target {
			// 配列の最後の値がtargetより大きい場合、配列の最後に追加するようにする
			// if nums[1] < target {
			// 	return index + 2
			// }
			return index + 1
		} else {
			return index
		}
	}

	// 配列の番分の位置の数値を確認
	var newNums []int
	half := len(nums) / 2
	if nums[half] == target {
		return index + half
	} else if nums[half] < target {
		newNums = nums[half:]
		index = index + half
	} else {
		newNums = nums[:half]
	}

	return searchInsertRecursive(newNums, target, index)
}

// @lc code=end
