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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n, m int
		fmt.Fscan(in, &n, &m)
		used := make([]int, n+1)
		deep := make([]int, n+1)
		num := make([][]int, n+1)
		rnum := make([][]int, n+1)
		for i := 1; i <= m; i++ {
			var x, y int
			fmt.Fscan(in, &x, &y)
			num[x] = append(num[x], y)
			rnum[y] = append(rnum[y], x)
		}
		var dfs1 func(int)
		dfs1 = func(x int) {
			used[x] = 1
			for _, i := range rnum[x] {
				if used[i] == 0 {
					dfs1(i)
				}
			}
		}
		mid := 0
		var dfs2 func(int)
		dfs2 = func(x int) {
			used[x]++
			for _, i := range num[x] {
				if used[i] == 1 {
					deep[i] = deep[x] + 1
					dfs2(i)
				} else if used[i] == 2 {
					mid = gcd(mid, deep[i]-deep[x]-1)
				}
			}
		}
		dfs1(1)
		dfs2(1)
		if mid != 0 {
			mid = abs(mid)
			for mid%5 == 0 {
				mid /= 5
			}
			for mid%2 == 0 {
				mid /= 2
			}
		}
		if mid == 1 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
