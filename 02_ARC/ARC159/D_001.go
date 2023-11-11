package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const INF = int(1e18)
const N = 200000
const N_ = 1 << 19
const MOD = 998244353

var Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var xx, xx_ [N * 2]int

func sort(ii []int, l, r int) {
	for l < r {
		i := l
		j := l
		k := r
		i_ := ii[l+Rand.Intn(MOD)%(r-l)]
		for j < k {
			if xx[ii[j]] == xx[i_] {
				j++
			} else if xx[ii[j]] < xx[i_] {
				ii[i], ii[j] = ii[j], ii[i]
				i++
				j++
			} else {
				k--
				ii[j], ii[k] = ii[k], ii[j]
			}
		}
		sort(ii, l, i)
		l = k
	}
}

var st [N_ * 2]int
var n_ int

func pul(i int) {
	st[i] = max(st[(i<<1)|0], st[(i<<1)|1])
}

func build(n int) {
	n_ = 1
	for n_ < n {
		n_ <<= 1
	}
	for i := 1; i < n_*2; i++ {
		st[i] = -INF
	}
}

func st_update(i, x int) {
	i += n_
	st[i] = max(st[i], x)
	for i > 1 {
		i >>= 1
		pul(i)
	}
}

func st_query(l, r int) int {
	x := -INF
	for l, r = l+n_, r+n_; l <= r; l, r = l>>1, r>>1 {
		if (l & 1) == 1 {
			x = max(x, st[l])
			l++
		}
		if (r & 1) == 0 {
			x = max(x, st[r])
			r--
		}
	}
	return x
}

var ft [N * 2]int

func ft_update(i, n, x int) {
	for i < n {
		ft[i] = max(ft[i], x)
		i |= i + 1
	}
}

func ft_query(i int) int {
	x := -INF
	for i >= 0 {
		x = max(x, ft[i])
		i &= i + 1
		i--
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var dp [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &xx[(i<<1)|0], &xx[(i<<1)|1])
		xx[(i<<1)|1]++
	}
	ii := make([]int, N*2)
	for i := 0; i < n*2; i++ {
		ii[i] = i
	}
	sort(ii, 0, n*2)
	m := 0
	var j int
	for i := 0; i < n*2; i = j {
		x := xx[ii[i]]
		xx_[m] = x
		j = i
		for j < n*2 && xx[ii[j]] == x {
			xx[ii[j]] = m
			j++
		}
		m++
	}
	for i := 0; i < n*2; i++ {
		ft[i] = -INF
	}
	build(m)
	ans := 0
	for i := 0; i < n; i++ {
		l := xx[(i<<1)|0]
		r := xx[(i<<1)|1]
		z1 := max(ft_query(l), 0) + xx_[r] - xx_[l]
		z2 := st_query(l, r)
		if z2 != -INF {
			z2 += xx_[r]
		}
		dp[i] = max(z1, z2)
		ans = max(ans, dp[i])
		ft_update(xx[i<<1|1], n*2, dp[i])
		st_update(xx[i<<1|1], dp[i]-xx_[r])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
