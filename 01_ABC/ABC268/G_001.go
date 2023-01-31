package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	s []byte
	x int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const mod = 998244353

	var n int
	fmt.Fscan(in, &n)
	s := make([]pair, n)
	for i := 0; i < n; i++ {
		var t string
		fmt.Fscan(in, &t)
		s[i].s = []byte(t)
		s[i].x = i
	}
	sort.Slice(s, func(i, j int) bool {
		return string(s[i].s) < string(s[j].s)
	})

	var a [500010]int
	for i := 0; i < n; i++ {
		a[s[i].x] += i + 1
	}
	for i := 0; i < n; i++ {
		for j, _ := range s[i].s {
			s[i].s[j] = 219 - s[i].s[j]
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return string(s[i].s) < string(s[j].s)
	})

	for i := 0; i < n; i++ {
		a[s[i].x] += i + 1
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, a[i]*499122177%mod)
	}
}
