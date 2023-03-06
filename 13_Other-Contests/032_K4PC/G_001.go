package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD = 1000000007

type P struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const SIZE = 100005

	var n int
	fmt.Fscan(in, &n)
	Us := 0
	Ds := 0
	A := make([]int, SIZE)
	down := make([]P, SIZE)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
		for Ds >= 2 && under(down[Ds-2], down[Ds-1], P{i, A[i]}) {
			Ds--
		}
		down[Ds] = P{i, A[i]}
		Ds++
	}
	B := make([]int, SIZE)
	up := make([]P, SIZE)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &B[i])
		for Us >= 2 && above(up[Us-2], up[Us-1], P{i, B[i]}) {
			Us--
		}
		up[Us] = P{i, B[i]}
		Us++
	}
	vx := make([]int, 0)
	S1 := make([]int, SIZE)
	for i := 0; i < Us-1; i++ {
		p := up[i]
		r := up[i+1]
		S1[i] = md(p.y*r.x-p.x*r.y, r.x-p.x)
		vx = append(vx, S1[i])
	}
	S2 := make([]int, SIZE)
	for i := 0; i < Ds-1; i++ {
		p := down[i]
		r := down[i+1]
		S2[i] = md(p.y*r.x-p.x*r.y+(r.x-p.x)-1, r.x-p.x)
		vx = append(vx, S2[i])
	}
	sort.Ints(vx)
	vx = unique(vx)
	now := 0
	left := make([]int, SIZE*2)
	for i := 0; i < len(vx); i++ {
		for now < Ds-1 && S2[now] <= vx[i] {
			now++
		}
		left[i] = now - 1
	}
	now = 0
	right := make([]int, SIZE*2)
	for i := len(vx) - 1; i >= 0; i-- {
		for now < Us-1 && S1[now] >= vx[i] {
			now++
		}
		right[i] = now - 1
	}
	ret := 0
	for i := 0; i+1 < len(vx); i++ {
		if left[i] == -1 || right[i+1] == -1 || !(vx[i]+1 < vx[i+1]) {
			continue
		}
		p := up[right[i+1]+1]
		q := down[left[i]+1]
		S := vx[i] + 1
		T := vx[i+1] - 1
		if p.x < q.x {
			lim := md(p.y*q.x-p.x*q.y, q.x-p.x)
			T = min(T, lim)
		} else if p.x > q.x {
			lim := md(q.y*p.x-q.x*p.y+(p.x-q.x)-1, p.x-q.x)
			S = max(S, lim)
		}
		if S > T {
			continue
		}
		ret += count(S, T, p.x, -1, p.y)
		ret += MOD - count(S, T, q.x, -1, q.y-1)
		ret %= MOD
	}
	for i := 0; i < len(vx); i++ {
		if left[i] == -1 || right[i] == -1 {
			continue
		}
		p := up[right[i]+1]
		q := down[left[i]+1]
		if (p.y-vx[i])*q.x < (q.y-vx[i])*p.x {
			continue
		}
		ret += count(vx[i], vx[i], p.x, -1, p.y)
		ret += MOD - count(vx[i], vx[i], q.x, -1, q.y-1)
		ret %= MOD
	}
	fmt.Println(ret)
}

func under(p, q, r P) bool {
	a := r.y - p.y
	b := -r.x + p.x
	c := p.y*r.x - p.x*r.y
	return a*q.x+b*q.y+c >= 0
}

func above(p, q, r P) bool {
	a := r.y - p.y
	b := -r.x + p.x
	c := p.y*r.x - p.x*r.y
	return a*q.x+b*q.y+c <= 0
}

func md(x, y int) int {
	if x >= 0 {
		return x / y
	}
	vl := md(-x, y)
	if (-x)%y == 0 {
		return -vl
	}
	return -vl - 1
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func count(s, t, n, a, b int) int {
	if a < 0 {
		return count(-t, -s, n, -a, b)
	}
	if md(a*s+b, n) == md(a*t+b, n) {
		return (t - s + 1) * get(md(a*s+b, n)) % MOD
	}
	if n <= a {
		sum := get(s+t) * get(t-s+1) % MOD
		for sum%2 != 0 {
			sum += MOD
		}
		sum /= 2
		return (count(s, t, n, a%n, b) + sum*(a/n)%MOD) % MOD
	}
	S := md(a*s+b, n)
	T := md(a*t+b, n)
	sum := count(S+1, T, a, n, -b-1)
	all := get(T*t - S*(s-1))
	return (all - sum + MOD) % MOD
}

func get(x int) int {
	if x < 0 {
		g := get(-x)
		if g == 0 {
			return 0
		}
		return MOD - g
	}
	return x % MOD
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
