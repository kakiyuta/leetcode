/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
package code

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	return recursiveMerge(list1, list2)
}

// 再起的にマージする
func recursiveMerge(current *ListNode, next *ListNode) *ListNode {
	var mergedList *ListNode
	var newCurrent *ListNode
	var newNext *ListNode
	if current.Val <= next.Val {
		mergedList = current
		newCurrent = next
		newNext = current
	} else {
		mergedList = next
		newCurrent = current
		newNext = next
	}

	if newNext.Next != nil {
		mergedList.Next = recursiveMerge(newCurrent, newNext.Next)
	} else {
		mergedList.Next = newCurrent
	}

	return mergedList
}

// @lc code=end
