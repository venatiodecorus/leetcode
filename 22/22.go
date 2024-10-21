package main

import (
	"fmt"
)

func main() {
	// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

	// Test cases
	fmt.Println("Should be [\"((()))\",\"(()())\",\"(())()\",\"()(())\",\"()()()\"]:", generateParenthesis(3))
	fmt.Println("Should be [\"()\"]:", generateParenthesis(1))

}

func generateParenthesis(n int) []string {
    // Repurpose backtrack solution from problem 17
	var result []string
    var backtrack func(index int, path string)
    backtrack = func(index int, path string) {
        if index == n*2 {
			if isValid(path) {
				result = append(result, path)
			}
            return
        }

        //letters := phoneLetters[rune(digits[index])]
		letters := []string{"(",")"}
        for _, letter := range letters {
            backtrack(index+1, path+letter)
        }
    }

    backtrack(0, "")
	return result
}


// Reuse valid parens check from problem 20
func isValid(s string) bool {
	open := []rune{'(','[','{'}
	close := []rune{')',']','}'}

	stack := Stack{}

	for _,c := range s {
		if contains(open, c) {
			stack.Push(c)
		}
		if contains(close, c) {
			if c == ')' && stack.Peek() == '(' {
				stack.Pop()
			}else if c == ']' && stack.Peek() == '[' {
				stack.Pop()
			}else if c == '}' && stack.Peek() == '{' {
				stack.Pop()
			} else {
				return false
			}
		}
	}

	if stack.IsEmpty() {
		return true
	} else {
		return false
	}
}

func contains(slice []rune, item rune) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}