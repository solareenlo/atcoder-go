package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var s = []string{}

func kurang(x int) bool {
	for i := 0; i < x; i++ {
		if s[i] == "1" {
			s[i] = "0"
			return true
		}
		s[i] = "1"
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var S string
	fmt.Fscan(in, &n, &S)
	S = reverseString(S)
	s = strings.Split(S, "")

	for len(s) < n {
		s = append(s, "0")
	}

	fmt.Fprint(out, 1)
	if kurang(n) {
		for i := n - 1; i >= 0; i-- {
			if s[i] == "0" {
				if kurang(i) {
					fmt.Fprint(out, 0)
				} else {
					break
				}
			} else {
				fmt.Fprint(out, 1)
			}
		}
	}
	fmt.Fprintln(out)
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
