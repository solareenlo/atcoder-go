package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type dat struct {
	x, y, d int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const P = 998244353
	const N = 305

	var n, d int
	fmt.Fscan(in, &n, &d)

	var pw2 [N]int
	pw2[0] = 1
	for i := 1; i <= n; i++ {
		pw2[i] = pw2[i-1] * 2 % P
	}

	a := make([]dat, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y)
		a[i].d = a[i].x - a[i].y
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].d < tmp[j].d
	})

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j || a[i].x < a[j].x && a[j].y < a[i].y && abs(a[i].d-a[j].d) <= d {
				p := make([]int, 0)
				for k := 1; k <= n; k++ {
					if k != i && k != j {
						if a[k].x < a[i].x || a[k].x > a[i].x+d || a[k].x == a[i].x && a[k].y < a[i].y {
							continue
						}
						if a[k].y < a[j].y || a[k].y > a[j].y+d || a[k].y == a[j].y && a[k].x < a[j].x {
							continue
						}
						if max(abs(a[k].d-a[i].d), abs(a[k].d-a[j].d)) <= d {
							p = append(p, a[k].d)
						}
					}
				}
				ans++
				l := 0
				for r := 0; r < len(p); r++ {
					for p[l] < p[r]-d {
						l++
					}
					ans += pw2[r-l]
				}
			}
		}
	}
	fmt.Println(ans % P)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
