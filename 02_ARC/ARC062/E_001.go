package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 505

var (
	c  = [N][4]int{}
	h  = [N]int{}
	v  = [4]int{}
	mp = map[int]int{}
)

func add(num, x int) {
	for i := 0; i < 4; i++ {
		mp[num] += x
		num = ((num & 1023) << 30) | (num >> 10)
	}
}

func Hash(a [4]int) int {
	res := 0
	for i := 0; i < 4; i++ {
		res |= (a[i] << (i * 10))
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fscan(in, &c[i][j])
		}
	}

	for i := 1; i <= n; i++ {
		h[i] = Hash(c[i])
		add(h[i], 1)
	}

	ans := 0
	for i := 1; i <= n-5; i++ {
		add(h[i], -1)
		for j := i + 1; j <= n; j++ {
			add(h[j], -1)
			for k := 0; k < 4; k++ {
				flg := true
				for l := 0; l < 4; l++ {
					tmp := [4]int{c[i][l], c[j][(3+k-l)%4], c[j][(6+k-l)%4], c[i][(l+1)%4]}
					v[l] = Hash(tmp)
					if _, ok := mp[v[l]]; !ok {
						flg = false
						break
					}
				}
				if !flg {
					continue
				}
				res := 1
				for l := 0; l < 4; l++ {
					res *= mp[v[l]]
					add(v[l], -1)
				}
				ans += res
				for l := 0; l < 4; l++ {
					add(v[l], 1)
				}
			}
			add(h[j], 1)
		}
	}
	fmt.Println(ans)
}
