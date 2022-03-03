package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	du := make([]int, n+1)
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		du[x]++
		du[y]++
	}
	for i := 1; i <= n; i++ {
		du[i]--
	}
	tmp := du[1 : n+1]
	sort.Ints(tmp)

	f := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = 1 << 60
	}
	f[0] = 0
	for i := 1; i <= n; i++ {
		j := i
		for j < n && du[j+1] == du[i] {
			j++
		}
		s := j - i + 1
		for k := 1; k <= s; k <<= 1 {
			w := k * du[i]
			for j := n - 2; j >= w; j-- {
				f[j] = min(f[j], f[j-w]+k)
			}
			s -= k
		}
		if s != 0 {
			w := s * du[i]
			for j := n - 2; j >= w; j-- {
				f[j] = min(f[j], f[j-w]+s)
			}
		}
		i = j
	}

	ans := 0
	for i := 0; i <= n-2; i++ {
		ans += max(0, n-f[n-2-i]-f[i]+1)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
