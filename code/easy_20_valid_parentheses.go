/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package code

import (
	"errors"
	"fmt"
)

/*
❌案1 : これでは要件を満たせられない
左から一文字ずつ確認し、かっこの始まりを検知したら同じ種類の閉じるを末尾まで検索する
最終的に閉じられてないものがあれば falseとなる

案2
左から一文字ずつ確認し、括弧の開始があればスタックに積み、括弧の閉じるを検知した時、スタックに積まれた始まりと種別が一致しているかで判定する
*/

// @lc code=start

// かっこの種別
// type BracketType int

// const (
// 	None          BracketType = iota // どの括弧にもマッチしない
// 	Parenthesis                      // ()
// 	CurlyBrace                       // {}
// 	SquareBracket                    // []
// )

// スタックを表す構造体
type Stack struct {
	elements []rune
}

// 新しいスタックを作成する関数
func NewStack() *Stack {
	return &Stack{elements: []rune{}}
}

// スタックに要素を追加する（push）
func (s *Stack) Push(element rune) {
	s.elements = append(s.elements, element)
}

// スタックから要素を取り出す（pop）
func (s *Stack) Pop() (rune, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("stack is empty")
	}
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element, nil
}

// スタックが空か判定
func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

// テスト対象
func isValid(s string) bool {
	runes := []rune(s)
	brackets := NewStack()
	for i := 0; i < len(runes); i++ {
		if isOpened(runes[i]) {
			brackets.Push(runes[i])
		} else if isClosed(runes[i]) {
			if brackets.isEmpty() {
				// 開始括弧が無い状態で閉じ括弧を検知した場合はその時点で不正となる
				return false
			}
			openingBracket, err := brackets.Pop()
			if err != nil {
				fmt.Println(err)
				return false
			}
			closer, err := getCloser(openingBracket)
			if err != nil {
				fmt.Println(err)
				return false
			}

			if runes[i] != closer {
				// 開始括弧と一致しない
				return false
			}
		}
	}

	return brackets.isEmpty()
}

// isOpen 開始括弧かを判定
func isOpened(r rune) bool {
	return r == '(' || r == '{' || r == '['
}

// isClosed 閉じ括弧かを判定
func isClosed(r rune) bool {
	return r == ')' || r == '}' || r == ']'
}

// 開始括弧に対応した閉じ括弧を取得
func getCloser(opening rune) (rune, error) {
	switch opening {
	case '(':
		return ')', nil
	case '{':
		return '}', nil
	case '[':
		return ']', nil
	default:
		return 0, errors.New("invalid opening bracket")
	}
}

// @lc code=end
