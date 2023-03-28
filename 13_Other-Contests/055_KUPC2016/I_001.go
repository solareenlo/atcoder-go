package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1e9 + 7

var num [2][100001]int
var dp [2][100001]int

func trim(a int) int {
	if a >= MOD {
		return a - MOD
	}
	return a
}

func dfs(g int, t int) {
	if dp[g][t] != -1 {
		return
	}
	if 2*g <= t {
		dfs(g, t-g)
		dfs(g+1, t-g)
		dp[g][t] = trim(dp[g][t-g] + dp[g+1][t-g])
		num[g][t] = trim(num[g][t-g] + num[g+1][t-g])
	} else {
		num[g][t] = 1
		dp[g][t] = t
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	pre, cur := 0, 1
	for g := 499; g > 0; g-- {
		pre, cur = cur, pre
		for t := 0; t <= 100000; t++ {
			if 2*g <= t {
				num[cur][t] = trim(num[cur][t-g] + num[pre][t-g])
				dp[cur][t] = trim(dp[cur][t-g] + dp[pre][t-g])
			} else {
				num[cur][t] = 1
				dp[cur][t] = t
			}
		}
	}
	var Q int
	fmt.Scan(&Q)
	for Q > 0 {
		Q--
		var N, C int
		fmt.Fscan(in, &N, &C)
		fmt.Fprintln(out, (dp[cur][N/C]*C+num[cur][N/C]*(N%C))%MOD)
	}
}
