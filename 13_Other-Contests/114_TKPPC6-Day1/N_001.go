package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(4e18)

	var n, w int
	fmt.Fscan(in, &n, &w)

	h := (n + w - 1) / w
	a := make([]int, n-1)
	b := make([]int, n-1)
	c := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
		a[i]++
		b[i]++
	}

	f := make([]int, (2*h)*(2*w))
	for i := range f {
		f[i] = INF
	}

	var get_min func(int, int, int, int) int
	get_min = func(li, ri, lj, rj int) int {
		res := INF
		for li, ri = li+h, ri+h; li < ri; li, ri = li>>1, ri>>1 {
			if (li & 1) != 0 {
				for l, r := lj+w, rj+w; l < r; l, r = l>>1, r>>1 {
					if (l & 1) != 0 {
						res = min(res, f[li*(2*w)+l])
						l++
					}
					if (r & 1) != 0 {
						r--
						res = min(res, f[li*(2*w)+r])
					}
				}
				li++
			}
			if (ri & 1) != 0 {
				ri--
				for l, r := lj+w, rj+w; l < r; l, r = l>>1, r>>1 {
					if (l & 1) != 0 {
						res = min(res, f[ri*(2*w)+l])
						l++
					}
					if (r & 1) != 0 {
						r--
						res = min(res, f[ri*(2*w)+r])
					}
				}
			}
		}
		return res
	}

	var update func(int, int, int)
	update = func(i, j, val int) {
		var Go func(int, int)
		Go = func(x, v int) {
			y := j + w
			f[x*(2*w)+y] = v
			for y = y >> 1; y > 0; y = y >> 1 {
				f[x*(2*w)+y] = min(f[x*(2*w)+(2*y)], f[x*(2*w)+(2*y+1)])
			}
		}
		Go(i+h, val)
		i += h
		for i = i >> 1; i > 0; i = i >> 1 {
			Go(i, min(f[(2*i)*(2*w)+(w+j)], f[(2*i+1)*(2*w)+(w+j)]))
		}
	}

	update((n-1)/w, (n-1)%w, 0)

	for v := n - 2; v >= 0; v-- {
		i := v / w
		j := v % w
		cur := INF
		if j+b[v] <= w {
			cur = min(cur, get_min(i, min(i+a[v], h), j, j+b[v]))
		} else {
			cur = min(cur, get_min(i, min(i+a[v], h), j, w))
			cur = min(cur, get_min(i+1, min(i+1+a[v], h), 0, j+b[v]-w))
		}
		cur += c[v]
		update(i, j, cur)
	}

	ans := get_min(0, 1, 0, 1)
	if ans < INF {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
