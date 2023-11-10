package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 300

var uu, vv [N][]int

func add(aa, cc []int, n, m int) bool {
	for i := 0; i < n; i++ {
		if aa[i] != 0 {
			if uu[i][i] != 0 {
				for j := 0; j < n; j++ {
					aa[j] ^= uu[i][j]
				}
				for h := 0; h < m; h++ {
					cc[h] ^= vv[i][h]
				}
			} else {
				copy(uu[i], aa)
				copy(vv[i], cc)
				return true
			}
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var bb [N][]int
	for i := 0; i < N; i++ {
		uu[i] = make([]int, N)
		vv[i] = make([]int, N)
		bb[i] = make([]int, N)
	}

	var m, n int
	fmt.Fscan(in, &m, &n)
	res := 0
	odd := false
	aa := make([]int, n)
	for h := 0; h < m; h++ {
		var s string
		fmt.Fscan(in, &s)
		for i := 0; i < n; i++ {
			bit, _ := strconv.Atoi(string(s[i]))
			aa[i] = bit
		}
		cc := make([]int, m)
		cc[h] = 1
		if add(aa, cc, n, m) {
			res++
		} else {
			x := 0
			for h_ := 0; h_ < m; h_++ {
				x ^= cc[h_]
			}
			if x != 0 {
				odd = true
			} else {
				for h_ := 0; h_ < m; h_++ {
					bb[h_][h] = cc[h_]
				}
			}
		}
	}
	for h := 0; h < m; h++ {
		uu[h] = make([]int, m)
	}
	if !odd {
		for h := 0; h < m; h++ {
			cc := make([]int, m)
			cc[h] = 1
			if !add(bb[h], cc, m, m) {
				for h_ := 0; h_ < m; h_++ {
					if cc[h_] == 0 {
						res--
						break
					}
				}
				break
			}
		}
	}
	fmt.Println(res)
}
