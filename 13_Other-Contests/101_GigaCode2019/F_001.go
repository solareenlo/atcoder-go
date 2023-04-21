package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var h, w int
	fmt.Fscan(in, &h, &w)
	var n int
	fmt.Fscan(in, &n)

	r := make([]int, n)
	c := make([]int, n)
	v := make([][]int, w)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i], &c[i])
		r[i]--
		c[i]--
		v[c[i]] = append(v[c[i]], r[i])
	}

	mi := make([]int, h)
	for i := 0; i < h; i++ {
		mi[i] = w
	}

	ans := 0
	for i := w - 1; i >= 0; i-- {
		used := make([]int, h+1)
		for j := 0; j < len(v[i]); j++ {
			mi[v[i][j]] = i
		}
		if i == 0 {
			for j := 0; j < h; j++ {
				used[j+1] = used[j] + 1
			}
		} else {
			for j := 0; j < len(v[i-1]); j++ {
				used[v[i-1][j]+1]++
			}
			for j := 0; j < h; j++ {
				used[j+1] += used[j]
			}
		}
		dq := make([]pair, 0)
		dq = append(dq, pair{-1, -1})
		for j := 0; j < h; j++ {
			c := mi[j]
			for dq[len(dq)-1].x > c {
				dq = dq[:len(dq)-1]
				nu := dq[len(dq)-1].y
				if used[j]-used[nu+1] > 0 {
					ans++
				}
			}
			for dq[len(dq)-1].x == c {
				dq = dq[:len(dq)-1]
			}
			dq = append(dq, pair{c, j})
		}
		c := -1
		for dq[len(dq)-1].x > c {
			y := dq[len(dq)-1].x
			dq = dq[:len(dq)-1]
			nu := dq[len(dq)-1].y
			if used[h]-used[nu+1] > 0 && y > i {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
