package main

import (
	"fmt"
)

func main() {
	fmt.Println("Should be 'PAHNAPLSIIGYIR': ", convert("PAYPALISHIRING",3))
	fmt.Println("Should be 'PINALSIGYAHRPI': ", convert("PAYPALISHIRING",4))
}

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    n := len(s)
    cycleLen := 2*numRows - 2
    res := make([]byte, n)
    idx := 0

    for i := 0; i < numRows; i++ {
        for j := 0; j+i < n; j += cycleLen {
            res[idx] = s[j+i]
            idx++
            if i != 0 && i != numRows-1 && j+cycleLen-i < n {
                res[idx] = s[j+cycleLen-i]
                idx++
            }
        }
    }

    return string(res)
}
