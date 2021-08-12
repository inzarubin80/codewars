package ValidBraces

import (
	"strings"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func ValidBraces(str string) bool {

	var expressionStack stack
	ar := strings.Split(str, "")
	st := ""
	for _, value := range ar {
		if value == "(" || value == "{" || value == "[" {
			expressionStack = expressionStack.Push(value)
		} else {

			if len(expressionStack) == 0 {
				return false
			}

			expressionStack, st = expressionStack.Pop()
			if !((value == ")" && st == "(") || (value == "]" && st == "[") || (value == "}" && st == "{")) {
				return false
			}

		}
	}

	return len(expressionStack) == 0

}

