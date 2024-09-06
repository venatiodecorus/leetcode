package main

import "fmt"

func main() {
	// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
	// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

	// test cases
	fmt.Println("Should be [ad,ae,af,bd,be,bf,cd,ce,cf]:", letterCombinations("23"))
	fmt.Println("Should be [a,b,c]:", letterCombinations("2"))
	fmt.Println("Should be []:", letterCombinations(""))
	fmt.Println("Should be [adg,adh,adi,aeg,aeh,aei,afg,afh,afi,bdg,bdh,bdi,beg,beh,bei,bfg,bfh,bfi,cdg,cdh,cdi,ceg,ceh,cei,cfg,cfh,cfi]:", letterCombinations("234"))
}

func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }

    phoneLetters := map[rune][]string{
        '2': {"a", "b", "c"},
        '3': {"d", "e", "f"},
        '4': {"g", "h", "i"},
        '5': {"j", "k", "l"},
        '6': {"m", "n", "o"},
        '7': {"p", "q", "r", "s"},
        '8': {"t", "u", "v"},
        '9': {"w", "x", "y", "z"},
    }

    var result []string
    var backtrack func(index int, path string)
    backtrack = func(index int, path string) {
        if index == len(digits) {
            result = append(result, path)
            return
        }

        letters := phoneLetters[rune(digits[index])]
        for _, letter := range letters {
            backtrack(index+1, path+letter)
        }
    }

    backtrack(0, "")
    return result
}