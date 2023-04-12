package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(2e17)

	var n int
	fmt.Fscan(in, &n)
	ab := make([]pair, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		ab[i] = pair{a + INF, b + INF}
	}
	ret := solve(ab, 70)
	fmt.Println(ret.x-INF, ret.y-INF)
}

func solve(ab []pair, ex int) pair {
	if ex == 0 {
		return pair{1, 0}
	}
	qab := make([][]pair, 2)
	qab[0] = make([]pair, 0)
	qab[1] = make([]pair, 0)
	for _, x := range ab {
		qab[x.y&1] = append(qab[x.y&1], pair{x.x / 2, x.y / 2})
		if x.x&1 != 0 {
			qab[(x.y+1)&1] = append(qab[(x.y+1)&1], pair{x.x / 2, (x.y + 1) / 2})
		}
	}
	c := make([]int, 2)
	for i := 0; i < 2; i++ {
		for _, x := range qab[i] {
			qc := (2*x.x + x.y) % 3
			c[i] ^= (qc + 1)
		}
	}
	for i := 0; i < 2; i++ {
		if c[1-i] == 0 {
			ans := solve(qab[i], ex-1)
			ans.x *= 2
			ans.y *= 2
			ans.y += i
			return ans
		}
	}
	for v := range ab {
		ab[v].x++
	}
	ans := solve(ab, ex)
	ans.x--
	return ans
}
