package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	r := make([]int, n)
	h := make([]int, n)
	s := make([]int, 100002)
	cnt := [100002][3]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i], &h[i])
		h[i]--
		s[r[i]]++
		cnt[r[i]][h[i]]++
	}

	for i := 0; i < 100001; i++ {
		s[i+1] += s[i]
	}

	for i := 0; i < n; i++ {
		win := s[r[i]-1] + cnt[r[i]][(h[i]+1)%3]
		even := cnt[r[i]][h[i]] - 1
		lose := n - 1 - win - even
		fmt.Println(win, lose, even)
	}
}
