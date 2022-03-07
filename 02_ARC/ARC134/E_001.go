package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const M = 998244353

var (
	n   int
	a   = [202]int{}
	ans int
)

func calc(SET []int) {
	f := [202][20]int{}
	cur := 0
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for cur < len(SET) && SET[cur] <= a[i] {
			cur++
		}
		for j := 0; j <= cur; j++ {
			f[i][j] += f[i-1][j] * j % M
			f[i][j] %= M
			f[i][j+1] += f[i-1][j] * (cur - j) % M
			f[i][j+1] %= M
		}
	}
	ans += f[n][len(SET)]
	ans %= M
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	tmp = make([]int, 1)
	tmp[0] = 1
	calc(tmp)
	tmp[0] = 2
	calc(tmp)
	tmp[0] = 4
	tmp = append(tmp, 8)
	calc(tmp)

	w := make([]int, 16)
	for i := 0; i < 16; i++ {
		w[i] = (i + 1) * 12
	}
	wn := [65540]int{}
	for i := 1; i < (1 << 16); i++ {
		v := make([]int, 0)
		for j := 0; j < 16; j++ {
			if (i>>j)&1 != 0 {
				v = append(v, w[j])
			}
		}
		wn[i] = 0
		for j := 1; j <= v[len(v)-1]; j++ {
			vv := make([]int, 0)
			for _, x := range v {
				if x%j > 0 {
					vv = append(vv, x%j)
				}
			}
			if len(vv) == 0 {
				continue
			}
			all1 := 1
			all2 := 1
			all12 := 1
			for _, x := range vv {
				tmp := 0
				if x == 1 {
					tmp = 1
				}
				all1 &= tmp
				tmp = 0
				if x == 2 {
					tmp = 1
				}
				all2 &= tmp
				tmp = 0
				if x == 1 || x == 2 {
					tmp = 1
				}
				all12 &= tmp
			}
			if all1 != 0 || all2 != 0 {
				wn[i] = 1
				break
			}
			ok0 := 0
			ok1 := 0
			ok2 := 1
			for _, x := range vv {
				ok0 |= (x & 1)
				tmp := 0
				if (x & 3) == 2 {
					tmp = 1
				}
				ok1 |= tmp
				tmp = 0
				if x%12 == 0 {
					tmp = 1
				}
				ok2 &= tmp
			}
			if ok0 != 0 || ok1 != 0 {
				continue
			}
			if ok2 != 0 {
				msk := 0
				for _, x := range vv {
					msk |= (1 << (x/12 - 1))
				}
				if wn[msk] == 0 {
					wn[i] |= 1
				} else {
					wn[i] |= 0
				}
			} else {
				oth := 1
				o4 := 0
				o8 := 0
				for _, x := range vv {
					tmp := 0
					if x == 4 {
						tmp = 1
					}
					o4 |= tmp
					tmp = 0
					if x == 8 {
						tmp = 1
					}
					o8 |= tmp
					tmp = 0
					if x == 4 || x == 8 {
						tmp = 1
					}
					oth &= tmp
				}
				if oth != 0 && o4 != 0 && o8 != 0 {
					wn[i] = 1
				}
			}
			if wn[i] != 0 {
				break
			}
		}
		if wn[i] == 0 {
			calc(v)
		}
	}
	sum := 1
	for i := 1; i <= n; i++ {
		sum = sum * a[i] % M
	}
	fmt.Println((sum - ans + M) % M)
}
