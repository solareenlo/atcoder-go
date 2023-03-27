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

	var N, Q int
	var x, v, t, l, r, ans, c, vc, p, vp, s [100009]int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &v[i])
	}
	t[Q] = 401
	fmt.Fscan(in, &Q)
	var cnt, vcnt int
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &t[i], &l[i], &r[i])
		r[i]++
		if t[i] <= 400 {
			c[t[i]+1]++
			cnt++
		} else if -250 <= v[i] && v[i] <= 250 {
			vc[v[i]+251]++
			vcnt++
		}
	}
	for i := 0; i <= 400; i++ {
		c[i+1] += c[i]
	}
	for i := 0; i <= 500; i++ {
		vc[i+1] += vc[i]
	}
	for i := 0; i < Q; i++ {
		if t[i] <= 400 {
			p[c[t[i]]] = i
			c[t[i]]++
		} else if -250 <= v[i] && v[i] <= 250 {
			vp[vc[v[i]+250]] = i
			vc[v[i]+250]++
		}
	}
	var ptr int
	for i := 0; i <= 400; i++ {
		for j := 0; j <= 100001; j++ {
			s[j] = 0
		}
		for j := 0; j < N; j++ {
			px := x[j] + i*v[j]
			if 0 <= px && px <= 100000 {
				s[px+1]++
			}
		}
		for j := 0; j <= 100000; j++ {
			s[j+1] += s[j]
		}
		for ptr < cnt && t[p[ptr]] == i {
			ans[p[ptr]] = s[r[p[ptr]]] - s[l[p[ptr]]]
			ptr++
		}
	}
	for i := -250; i <= 250; i++ {
		for j := 0; j <= 100001; j++ {
			s[j] = 0
		}
		for j := 0; j < N; j++ {
			if v[j] == i {
				s[x[j]+1]++
			}
		}
		for j := 0; j <= 100000; j++ {
			s[j+1] += s[j]
		}
		for j := 0; j < Q; j++ {
			if t[j] > 400 {
				pl := l[j] - i*t[j]
				if pl < 0 {
					pl = 0
				} else if pl > 100001 {
					pl = 100001
				}
				pr := r[j] - i*t[j]
				if pr < 0 {
					pr = 0
				} else if pr > 100001 {
					pr = 100001
				}
				ans[j] += s[pr] - s[pl]
			}
		}
	}
	for i := 0; i < Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
