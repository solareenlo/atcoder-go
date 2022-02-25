package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for j := 0; j < t; j++ {
		var s string
		fmt.Fscan(in, &s)
		N := len(s)
		A := strings.Count(s, "a")
		B := N - A
		a := 0
		b := 0
		if A%2 == 1 && s[max(N-3, 0):] == "bbb" {
			if strings.Contains(s, "ba") {
				a--
				b = 2
			}
			if strings.Contains(s, "baa") {
				a--
			}
		}
		if s[len(s)-1] == 'a' || b > 0 {
			for i := 0; i < N; i++ {
				if i+2 < len(s) && s[i:i+2] == "ab" {
					a++
				}
				if i+3 < len(s) && s[i:i+3] == "aab" {
					a++
				}
			}
			a += a % 2
			fmt.Fprintln(out, strings.Repeat("b", B-b)+strings.Repeat("a", A-a))
		} else {
			re := regexp.MustCompile("a(b*)a")
			fmt.Fprintln(out, re.ReplaceAllString(s, "$1"))
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
