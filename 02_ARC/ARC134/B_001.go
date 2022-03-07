package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var S string
	fmt.Fscan(in, &n, &S)
	s := strings.Split(S, "")

	ls := 0
	rs := n - 1
	ch := "abcdefghijklmnopqrstuvwxyz"
	for _, t := range ch {
		l := ls
		r := rs
		for r > l {
			for s[l][0] <= byte(t) && l < r {
				l++
			}
			if s[r][0] == byte(t) && s[r] < s[l] {
				s[l], s[r] = s[r], s[l]
				l++
				rs = r - 1
				ls = l
			}
			r--
		}
	}
	fmt.Fprintln(out, strings.Join(s, ""))
}
