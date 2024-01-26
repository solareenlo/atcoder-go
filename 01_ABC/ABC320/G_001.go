package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 110
const M = 1e5 + 10

var vec [10][N][]int
var match map[int]int
var n, m, tp int
var vis [N * M]bool
var stk [N * M]int
var s [N]string

func Dfs(u, num, tim int) int {
	for _, t := range vec[num][u] {
		for v := t; v <= tim; v += m {
			if !vis[v] {
				vis[v] = true
				tp++
				stk[tp] = v
				if match[v] == 0 || Dfs(match[v], num, tim) != 0 {
					match[v] = u
					return 1
				}
			}
		}
	}
	return 0
}

func check(num, tim int) bool {
	res := 0
	match = make(map[int]int)
	for i := 1; i <= n; i++ {
		for tp > 0 {
			vis[stk[tp]] = false
			tp--
		}
		res += Dfs(i, num, tim)
	}
	return res == n
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		for j := 0; j < m; j++ {
			if len(vec[s[i][j]-'0'][i]) < n {
				vec[s[i][j]-'0'][i] = append(vec[s[i][j]-'0'][i], j)
			}
		}
	}
	res := INF
	for i := 0; i < 10; i++ {
		l := 0
		r := n * m
		for l <= r {
			mid := (l + r) >> 1
			if check(i, mid) {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		res = min(res, l)
	}
	if res > n*m {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
