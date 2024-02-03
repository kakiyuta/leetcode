/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */
package code

import "strings"

// @lc code=start
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// @lc code=end
