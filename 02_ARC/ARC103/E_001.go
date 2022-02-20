package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	s = "0" + s
	t := reverseString(s)

	if s != t || s[1] == 48 {
		fmt.Fprintln(out, -1)
		return
	}

	p := 1
	for i := 2; i < len(s); i++ {
		fmt.Fprintln(out, p, i)
		if s[i-1] != 48 {
			p = i
		}
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
