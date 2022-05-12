package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, m int
	ans  int
	mx   = make([]int, 200200)
	mi   = make([]int, 200200)
	a    = make([]int, 200200)
	cnt  = make([]int, 400400)
)

func solve(l, r int) {
	if l == r {
		ans++
		return
	}
	M := (l + r) >> 1
	solve(l, M)
	solve(M+1, r)
	mx[M] = a[M]
	mi[M] = a[M]
	mx[M+1] = a[M+1]
	mi[M+1] = a[M+1]
	for i := M + 2; i <= r; i++ {
		mx[i] = max(mx[i-1], a[i])
		mi[i] = min(mi[i-1], a[i])
	}
	for i := M - 1; i >= l; i-- {
		mx[i] = max(mx[i+1], a[i])
		mi[i] = min(mi[i+1], a[i])
	}
	for L := l; L <= M; L++ {
		for k := 0; k <= m; k++ {
			R := mx[L] - mi[L] + L - k
			if R > M && R <= r && mx[R] < mx[L] && mi[R] > mi[L] {
				ans++
			}
		}
	}
	for R := M + 1; R <= r; R++ {
		for k := 0; k <= m; k++ {
			L := R - mx[R] + mi[R] + k
			if L >= l && L <= M && mx[R] > mx[L] && mi[R] < mi[L] {
				ans++
			}
		}
	}
	for L, R, rs := M, M+1, M+1; L >= l; L-- {
		for R <= r && mx[R] < mx[L] {
			cnt[R+mi[R]]++
			R++
		}
		for rs < R && mi[rs] > mi[L] {
			cnt[rs+mi[rs]]--
			rs++
		}
		for k := 0; k <= m; k++ {
			if L+mx[L] > k {
				ans += cnt[L+mx[L]-k]
			}
		}
	}
	for R := M + 1; R <= r; R++ {
		cnt[R+mi[R]] = 0
	}
	for L, R, ls := M, M+1, M; R <= r; R++ {
		for L >= l && mx[L] < mx[R] {
			cnt[L-mi[L]+n]++
			L--
		}
		for ls > L && mi[ls] > mi[R] {
			cnt[ls-mi[ls]+n]--
			ls--
		}
		for k := 0; k <= m; k++ {
			ans += cnt[R-mx[R]+k+n]
		}
	}
	for L := M; L >= l; L-- {
		cnt[L-mi[L]+n] = 0
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	solve(1, n)

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
