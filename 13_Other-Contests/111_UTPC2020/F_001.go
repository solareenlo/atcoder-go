package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100000

	var n, k, x, y int
	fmt.Fscan(in, &n, &k, &x, &y)
	s := 0
	var cnt [N + 5]int
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		cnt[a]++
		s += a
	}
	s = (s + k - 1) / k
	ans := 1 << 60
	tmp := 0
	for i := N; i > k; i-- {
		ans = min(ans, tmp+max(s, i)*y)
		if s > i && i+cnt[i] > s {
			ans = min(ans, tmp+(s-i)*x+i*y)
		}
		tmp += cnt[i] * x
		s -= cnt[i]
		cnt[i-k] += cnt[i]
	}
	s = 0
	for i := 1; i <= k; i++ {
		s += i * cnt[i]
	}
	for i := k; i > 0; i-- {
		for j := 0; j < cnt[i]; j++ {
			ans = min(ans, tmp+max(i, (s+k-1)/k)*y)
			s -= i
			tmp += x
		}
	}
	ans = min(ans, tmp)
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
