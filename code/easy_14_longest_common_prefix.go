/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package code

// @lc code=start
// longestCommonPrefix 文字列配列ないで共通している接頭辞を返す
// 文字共通する接頭辞がない場合は空配列を返す
/*
	案1
	文字列配列を文字配列の配列に変換
	[
		"abcd",
		"efgh"
	]
	↓
	[
		['a', 'b', 'c', 'd'],
		['e', 'f', 'g', 'h']
	]
	そして、配列の先頭から一つづく確認してって、共通している接頭辞を算出する
	途中で一致しないケースがでた場合はその時で処理を中断する
	一文字目で一致しない、または対象の配列内でindex範囲外になった場合も同様
*/
func longestCommonPrefix(strs []string) string {

	count := len(strs)
	if count == 0 {
		return ""
	} else if count == 1 {
		// 要素数が一つの時はその一つが共通の接頭辞となる
		return strs[0]
	}
	base := []rune(strs[0])

	var searchTargets [][]rune = make([][]rune, count-1)
	for i := 1; i < len(strs); i++ {
		searchTargets[i-1] = []rune(strs[i])
	}

	var result []rune
OuterLoop:
	for i, char := range base {
		for _, target := range searchTargets {
			if i >= len(target) {
				break OuterLoop
			}
			if char != target[i] {
				break OuterLoop
			}
		}

		result = append(result, char)
	}

	return string(result)

}

// @lc code=end
