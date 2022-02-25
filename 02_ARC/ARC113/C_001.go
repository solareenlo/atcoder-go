package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	n := len(s)
	s += " "
	cnt := [128]int{}
	res := 0
	for i := n; i > 0; i-- {
		if s[i-1] == s[i] {
			for c := 'a'; c <= 'z'; c++ {
				if s[i] != byte(c) {
					res += cnt[c]
					cnt[s[i]] += cnt[c]
					cnt[c] = 0
				}
			}
		}
		cnt[s[i]]++
	}
	fmt.Println(res)
}
