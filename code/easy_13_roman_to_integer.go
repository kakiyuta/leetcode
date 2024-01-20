/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package code

/*
	案1
	一文字ループで確認し、各文字に対応した数値で加算していく。
	ただし、
	・I
	・X
	・C
	については次の文字によって挙動を変える
	・I：次がVなら4、Xなら9
	・X：次がLなら40、Cなら90
	・C：次がDなら400、Mなら900
*/

// @lc code=start
func romanToInt(s string) int {
	result := 0

	// 文字列を各文字に分解してループで処理していく
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		num := romaCharToInt(runes[i])
		next := i + 1
		if next >= len(runes) {
			result += num
			continue
		}

		// ローマ数字は基本的に、右から左で数値が小さくなっていくが、もし次の文字が今の文字より大きい場合は減算する
		nextNum := romaCharToInt(runes[next])
		if num < nextNum {
			result += (nextNum - num)
			i++ // nextの文字は減算で使ったため、nextの評価を飛ばす
		} else {
			result += num
		}

	}
	return result
}

// romaCharToInt ローマ数字の文字を数値に変換
func romaCharToInt(c rune) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}

	return 0
}

// @lc code=end

/*
M 1000
C 100
M 1000
X 10
C 100
I 1
V 5

2216
1994


*/
