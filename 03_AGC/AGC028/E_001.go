package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N int
	f = make([][2]int, 600100)
)

func upDate(x, v0, v1 int) {
	x += N
	f[x][0] = v0
	f[x][1] = v1
	for x >>= 1; x > 0; x >>= 1 {
		f[x][0] = max(f[x*2][0], f[x*2+1][0])
		f[x][1] = max(f[x*2][1], f[x*2+1][1])
	}
}

var ans0, ans1 int

func query(l, r int) int {
	ans0 = 0
	ans1 = -(1 << 30)
	l += N - 1
	r += N + 1
	for ; l != (r ^ 1); l, r = l>>1, r>>1 {
		if l&1 == 0 {
			ans0 = max(ans0, f[l^1][0])
			ans1 = max(ans1, f[l^1][1])
		}
		if r&1 != 0 {
			ans0 = max(ans0, f[r^1][0])
			ans1 = max(ans1, f[r^1][1])
		}
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for N = 1; N <= n; N <<= 1 {
	}

	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 1; i <= 2*N; i++ {
		f[i][1] = -(1 << 30)
	}

	mx := 0
	q := make([]int, n+2)
	for i := 1; i <= n; i++ {
		if p[i] > mx {
			mx = p[i]
			q[i] = 1
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := n; i > 0; i-- {
		query(p[i]+1, n)
		if q[i] != 0 {
			dp[i][0] = ans0 + 2
			dp[i][1] = ans1 + 2
		} else {
			dp[i][0] = ans1 + 1
			dp[i][1] = ans0 + 1
		}
		upDate(p[i], dp[i][0], dp[i][1])
		q[i] += q[i+1]
	}

	query(1, n)
	if q[1]&1 != 0 {
		if ans1 < q[1] {
			fmt.Println(-1)
			return
		}
	} else if ans0 < q[0] {
		fmt.Println(-1)
		return
	}

	ca, cb, ma, mb := 0, 0, 0, 0
	for i := 1; i <= n; i++ {
		upDate(p[i], 0, -(1 << 30))
		ca1 := ca
		cb1 := cb
		ma1 := ma
		mb1 := mb
		if p[i] > ma1 {
			ma1 = p[i]
			ca1++
		}
		if q[i+1]+cb1-ca1 >= 0 {
			query(ma1+1, n)
			if (q[i+1]+cb1-ca1)&1 != 0 {
				if ans1 >= q[i+1]+cb1-ca1 {
					fmt.Print(0)
					if p[i] > ma {
						ma = p[i]
						ca++
					}
					continue
				}
			} else {
				if ans0 >= q[i+1]+cb1-ca1 {
					fmt.Print(0)
					if p[i] > ma {
						ma = p[i]
						ca++
					}
					continue
				}
			}
		}
		if q[i+1]+ca1-cb1 >= 0 {
			query(mb1+1, n)
			if (q[i+1]+ca1-cb1)&1 != 0 {
				if ans1 >= q[i+1]+ca1-cb1 {
					fmt.Print(0)
					if p[i] > ma {
						ma = p[i]
						ca++
					}
					continue
				}
			} else {
				if ans0 >= q[i+1]+ca1-cb1 {
					fmt.Print(0)
					if p[i] > ma {
						ma = p[i]
						ca++
					}
					continue
				}
			}
		}
		fmt.Print(1)
		if p[i] > mb {
			mb = p[i]
			cb++
		}
		continue
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
