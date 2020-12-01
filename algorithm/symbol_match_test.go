package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

func TestMatchSymbol(t *testing.T) {
	fmt.Println(matchSymbol("()[]{}"))
}

func matchSymbol(str string) bool {
	stack := list.New()
	for _, ch := range str {
		if ch == '{' || ch == '[' || ch == '(' {
			stack.PushBack(ch)
			continue
		}
		popCh := pop(stack)
		if popCh != leftOf(ch) {
			return false
		}
	}
	return true
}

func leftOf(ch rune) rune {
	if ch == '}' {
		return '{'
	}
	if ch == ']' {
		return '['
	}
	return '('
}

func pop(stack *list.List) interface{} {
	e := stack.Back()
	if e != nil {
		stack.Remove(e)
		return e.Value
	}
	return nil
}
