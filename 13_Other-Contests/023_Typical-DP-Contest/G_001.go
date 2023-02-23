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

	const N = 1000005
	const INF = int(2e18)

	var str string
	fmt.Fscan(in, &str)
	n := len(str)
	var m int
	fmt.Fscan(in, &m)
	dp := make([]int, N)
	dp[n] = 1
	dp[n+1] = 0
	last := make([]int, N)
	for i := 0; i < 26; i++ {
		last[i] = n
	}
	cl := make([][]int, 26)
	for i := n - 1; i >= 0; i-- {
		last[str[i]-'a'] = i
		dp[i] = 1
		for j := 0; j < 26; j++ {
			dp[i] = min(INF, dp[i]+dp[last[j]+1])
		}
	}
	for i := 0; i < n; i++ {
		cl[str[i]-'a'] = append(cl[str[i]-'a'], i)
	}
	for i := 0; i < 26; i++ {
		cl[i] = append(cl[i], n)
	}
	m++
	if dp[0] < m {
		fmt.Fprintln(out, "Eel")
		return
	}
	now := -1
	p := make([]int, 26)
	for now < n {
		if m == 1 {
			fmt.Fprintln(out)
			break
		} else {
			m--
		}
		for i := 0; i < 26; i++ {
			for now >= cl[i][p[i]] {
				p[i]++
			}
			if m <= dp[cl[i][p[i]]+1] {
				fmt.Fprintf(out, "%c", 'a'+i)
				now = cl[i][p[i]]
				break
			} else {
				m -= dp[cl[i][p[i]]+1]
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
