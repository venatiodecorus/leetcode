package main

import (
	"fmt"
)

func main() {
	// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

	// An input string is valid if:

    // Open brackets must be closed by the same type of brackets.
    // Open brackets must be closed in the correct order.
    // Every close bracket has a corresponding open bracket of the same type.

	// Test cases
	fmt.Println("Should be true:", isValid("()"))
	fmt.Println("Should be true:", isValid("()[]{}"))
	fmt.Println("Should be false:", isValid("(]"))
	fmt.Println("Should be false:", isValid("([)]"))
	fmt.Println("Should be true:", isValid("{[]}"))
	fmt.Println("Should be false:", isValid("]"))
}

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