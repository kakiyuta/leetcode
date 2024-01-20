/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
/*
	案1：一度文字列に変換して、さらに一文字の文字配列に分解する、そして、配列の中身を反転させて比較する
	・文字列に変換する方法
	strconv.FormatFloat(10000000000000000000000000, 'f', -1, 64)
	・文字列を反転する方法
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)


	案2：桁数を確認後、各桁の数値ごとを配列に入れて反転させて比較する

	案3：桁数を確認後、各桁の数値を抽出（例：987 / 100で9を抽出。そして900を引く）していく

	メモ：
	桁数は一度文字列にしてから文字列の長さで求めることができる。
	→一度文字列にするなら案1でいいんじゃない？
*/

package code

import (
	"strconv"
)

// @lc code=start
func isPalindrome(x int) bool {
	// 案1で実装
	str := strconv.Itoa(x)

	// 文字列を逆転
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reverses := string(runes)

	return str == reverses
}

// @lc code=end
