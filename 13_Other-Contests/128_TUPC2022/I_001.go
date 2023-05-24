package main

import (
	"fmt"
	"sort"
)

const maxn = 10000010
const mod = 998244353

var cur int
var pri, mu, sum_mu [maxn]int
var vis [maxn]bool
var mp_mu map[int]int

func main() {
	mu[1] = 1
	for i := 2; i < maxn; i++ {
		if !vis[i] {
			cur++
			pri[cur] = i
			mu[i] = -1
		}
		for j := 1; j <= cur && i*pri[j] < maxn; j++ {
			vis[i*pri[j]] = true
			if i%pri[j] != 0 {
				mu[i*pri[j]] = -mu[i]
			} else {
				mu[i*pri[j]] = 0
				break
			}
		}
	}
	for i := 1; i < maxn; i++ {
		sum_mu[i] = sum_mu[i-1] + mu[i]
	}

	var l, r int
	fmt.Scan(&l, &r)

	v := make([]int, 0)
	for i := 1; i*i <= r; i++ {
		v = append(v, i)
		v = append(v, r/i)
	}
	if l > 1 {
		for i := 1; i*i <= (l - 1); i++ {
			v = append(v, i)
			v = append(v, (l-1)/i)
		}
	}
	sort.Ints(v)
	v = unique(v)

	last := 0
	ans := 0
	mp_mu = make(map[int]int)
	for _, x := range v {
		c := r/x - (l-1)/x
		ans += (powMod(2, c) - 1) * ((S_mu(x) - S_mu(last)) % mod)
		ans %= mod
		last = x
	}

	ans %= mod
	if ans < 0 {
		ans += mod
	}
	fmt.Println(ans)
}

func S_mu(x int) int {
	if x < maxn {
		return sum_mu[x]
	}
	if _, ok := mp_mu[x]; ok {
		return mp_mu[x]
	}
	ret := 1
	for i, j := 2, 0; i <= x; i = j + 1 {
		j = x / (x / i)
		ret -= S_mu(x/i) * (j - i + 1)
	}
	mp_mu[x] = ret
	return mp_mu[x]
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
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
