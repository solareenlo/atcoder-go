package main

import (
	"bufio"
	"fmt"
	"os"
)

const MX = 1000010

var nn int
var nxt [MX][26]int
var nval [MX]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var hay string
	fmt.Fscan(in, &hay)
	hay = " " + hay + " "
	n := len(hay) - 1
	var m int
	fmt.Fscan(in, &m)
	var w [5010]string
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &w[i])
	}
	nn = 1
	for i := 0; i < m; i++ {
		var val int
		fmt.Fscan(in, &val)
		put(w[i], val)
	}

	var dp [200010]int
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		p := 1
		for l := 1; l <= 200 && i-l >= 0; l++ {
			c := hay[i-l+1] - 'a'
			if c > 26 || nxt[p][c] == 0 {
				break
			}
			p = nxt[p][c]
			dp[i] = max(dp[i], dp[i-l]+nval[p])
		}
	}
	fmt.Println(dp[n])
}

func put(pt string, x int) {
	n := len(pt)
	p := 1
	for i := n - 1; 0 <= i; i-- {
		d := pt[i] - 'a'
		if nxt[p][d] == 0 {
			nn++
			nxt[p][d] = nn
		}
		p = nxt[p][d]
	}
	nval[p] = max(nval[p], x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
