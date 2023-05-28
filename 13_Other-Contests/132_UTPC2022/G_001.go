package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	var s [200020]int
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + a[i]
	}
	ans := -(1 << 60)
	p := n/2 - 1
	if n%2 != 0 {
		p = (n + 1) / 2
	}
	r := p / 3
	f := 1
	if n%2 != 0 {
		f = 0
	}
	p -= r * 3
	f += r * 2
	for p <= n {
		t := s[n] - s[n-p]*2 + f*k
		ans = max(ans, t)
		p += 3
		f -= 2
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
