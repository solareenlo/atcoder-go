package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	mp := make(map[pair]int)
	ans := 0
	sum := 0.0
	sum1 := 0
	sum2 := 0
	mp[pair{0, 0}]++
	for i := 0; i < n; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		sum += float64(p) / float64(q)
		sum1 = (sum1 + p*invMod(q, mod1)%mod1) % mod1
		sum2 = (sum2 + p*invMod(q, mod2)%mod2) % mod2
		v1 := (sum1 - int(math.Floor(sum)) + mod1) % mod1
		v2 := (sum2 - int(math.Floor(sum)) + mod2) % mod2
		ans += mp[pair{v1, v2}]
		ans += mp[pair{(v1 - 1), (v2 - 1)}]
		ans += mp[pair{(v1 + 1), (v2 + 1)}]
		mp[pair{v1, v2}]++
	}
	fmt.Println(ans)
}

const mod1 = 897581057
const mod2 = 880803841

func PowMod(a, n, mod int) int {
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

func invMod(a, mod int) int {
	return PowMod(a, mod-2, mod)
}
