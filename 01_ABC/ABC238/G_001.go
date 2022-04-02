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

	const N = 1000001
	v := make([]int, N)
	v[0] = 1
	for i := 1; i < N; i++ {
		v[i] = v[i-1] * 23
	}
	vis := make([]bool, N+1)
	minp := make([]int, N+1)
	t := make([]int, N)
	top := 0
	for i := 2; i < N; i++ {
		if !vis[i] {
			top++
			t[top] = i
			minp[i] = top
		}
		for j := 1; j <= top && t[j]*i <= N; j++ {
			vis[t[j]*i] = true
			minp[t[j]*i] = j
			if i%t[j] == 0 {
				break
			}
		}
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	val := make([]int, N)
	cnt := make([]int, N)
	val[0] = 1
	sum := 1
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for x > 1 {
			cur := minp[x]
			x /= t[cur]
			sum -= v[cur] * cnt[cur]
			cnt[cur] = (cnt[cur] + 1) % 3
			sum += v[cur] * cnt[cur]
		}
		val[i] = sum
	}
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		if val[l] == val[r] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
