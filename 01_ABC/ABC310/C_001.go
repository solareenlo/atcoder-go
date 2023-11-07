package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	t := make(map[string]bool)

	var n int
	fmt.Fscan(in, &n)
	for n > 0 {
		n--
		var s string
		fmt.Fscan(in, &s)
		ss := s
		ss = reverseString(ss)
		t[min(s, ss)] = true
	}
	fmt.Println(len(t))
}

func min(a, b string) string {
	if len(a) < len(b) || (len(a) == len(b)) && a < b {
		return a
	}
	return b
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
