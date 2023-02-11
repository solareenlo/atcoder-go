package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	divide := [4]string{"dream", "dreamer", "erase", "eraser"}

	var s string
	fmt.Fscan(in, &s)
	s = reverseString(s)

	for i := 0; i < 4; i++ {
		divide[i] = reverseString(divide[i])
	}

	ok := true
	for i := 0; i < len(s); i++ {
		ok2 := false
		for j := 0; j < 4; j++ {
			div := divide[j]
			if i+len(div) <= len(s) && s[i:i+len(div)] == div {
				ok2 = true
				i += len(div) - 1
			}
		}
		if !ok2 {
			ok = false
			break
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
