package main

import (
	"fmt"
	"math"
)

func get(r float64, n int) int {
	if r <= 0 {
		return 0
	}
	cnt := make([]int, n+2)
	for i := 2; i <= n+1; i++ {
		cnt[i] = int(r * float64(i))
		cnt[i] = int(math.Min(float64(cnt[i]), float64(i-1)))
	}
	for i := 1; i <= n+1; i++ {
		for j := 2 * i; j <= n+1; j += i {
			cnt[j] -= cnt[i]
		}
	}
	ans := 0
	for _, c := range cnt {
		ans += c
	}
	return ans
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for ; m > 0; m-- {
		var p, q float64
		var k int
		fmt.Scan(&p, &q, &k)
		lo := 0.0
		hi := 1.0
		dis := 0.0
		for {
			d := (lo + hi) / 2.0
			tmp1 := get(p/q+d, n)
			tmp2 := get(p/q-d, n)
			cnt := tmp1 - tmp2
			if cnt == k {
				dis = d
				break
			}
			if cnt < k {
				lo = d
			} else {
				hi = d
			}
		}

		l := math.Max(p/q-dis, 1e-9)
		r := math.Min(p/q+dis, 1-1e-9)
		nl, dl := int(p), int(q)
		nr, dr := int(p), int(q)
		for i := 2; i <= n+1; i++ {
			minn := int(math.Ceil(l*float64(i) - 1e-9))
			maxn := int(r*float64(i) + 1e-9)
			if float64(minn)/float64(i) < float64(nl)/float64(dl) {
				nl = minn
				dl = i
			}
			if float64(maxn)/float64(i) > float64(nr)/float64(dr) {
				nr = maxn
				dr = i
			}
		}
		gl := gcd(int(nl), int(dl))
		nl /= gl
		dl /= gl
		gr := gcd(int(nr), int(dr))
		nr /= gr
		dr /= gr
		sm := (nl*dr + nr*dl) * int(q)
		md := 2 * int(p) * dl * dr
		if sm < md {
			fmt.Println(dl - 1)
		} else if sm > md {
			fmt.Println(dr - 1)
		} else {
			fmt.Println(min(dl, dr) - 1)
		}
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
