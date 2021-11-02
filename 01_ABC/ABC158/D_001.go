package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s1, s2 string
	var n int
	fmt.Fscan(in, &s1, &n)

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			s1, s2 = s2, s1
		} else {
			var f int
			var c string
			fmt.Fscan(in, &f, &c)
			if f == 1 {
				s2 += c
			} else {
				s1 += c
			}

		}
	}

	fmt.Println(reverseString(s2) + s1)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
