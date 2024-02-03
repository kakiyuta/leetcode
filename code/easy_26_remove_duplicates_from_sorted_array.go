/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */
package code

/*
メモ
numsを別のスライスに置き換えると、アドレスが変わり呼び出し元で同じスライスを参照できなくなる。
❌ nums = []int{1,2}
問題の使用上
nums[0] = 1
nums[1] = 2
といって方法で値を変えていく必要がある
*/

// @lc code=start
func removeDuplicates(nums []int) int {
	uniqueNums := uniqueInts(nums)

	// 問題の仕様上、numsはソート済みのため、uniqueNumsを改めてソートしたりしない
	// ただ、実際のプロダクトの処理的にはソートした方が良さそう

	for i, num := range uniqueNums {
		nums[i] = num
	}

	return len(uniqueNums)

}

// uniqueInts スライス内の整数をユニークな値にする
func uniqueInts(nums []int) []int {
	seen := make(map[int]bool)
	var unique []int

	for _, v := range nums {
		if _, ok := seen[v]; !ok {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

// @lc code=end
