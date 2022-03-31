package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	type tuple struct {
		x, y, i int
		d       float64
	}
	data := make([]tuple, n)
	const theta = 1.41421356
	ct := math.Cos(theta)
	st := math.Sin(theta)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		data[i].x = x
		data[i].y = y
		data[i].i = i + 1
		data[i].d = float64(x)*st + float64(y)*ct
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].d < data[j].d
	})

	const EPS = 1e-6
	kk := float64(k) * (1.0 + EPS)
	l := 0
	cnt := 0
	const N = 200005
	ans := make([]tuple, 2*N)
	for i := 0; i < n; i++ {
		for data[i].d-data[l].d > kk {
			l++
		}
		for j := l; j < i; j++ {
			if i == j {
				continue
			}
			xx := data[i].x - data[j].x
			yy := data[i].y - data[j].y
			if xx*xx+yy*yy <= k*k {
				ans[cnt].x = min(data[i].i, data[j].i)
				ans[cnt].y = max(data[i].i, data[j].i)
				cnt++
			}
		}
	}
	tmp := ans[:cnt]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].x < tmp[j].x || (tmp[i].x == tmp[j].x && tmp[i].y < tmp[j].y)
	})

	fmt.Println(cnt)
	for i := 0; i < cnt; i++ {
		fmt.Println(ans[i].x, ans[i].y)
	}
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
