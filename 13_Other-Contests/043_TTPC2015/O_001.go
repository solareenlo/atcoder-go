package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

var N int
var a []int
var lu, ru, ld, rd []int
var LU, RU, LD, RD []int
var vc [100001][]int
var bit, bit2 [100001]int

type P struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	sum = 1
	lu = make([]int, 100000)
	ru = make([]int, 100000)
	LU = make([]int, 100000)
	RU = make([]int, 100000)
	ld = make([]int, 100000)
	rd = make([]int, 100000)
	LD = make([]int, 100000)
	RD = make([]int, 100000)

	fmt.Fscan(in, &N)
	a = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	calc(lu, LU)
	I++
	for i := 0; i < N; i++ {
		a[i] = N - 1 - a[i]
	}
	calc(ld, LD)
	I++
	for i := 0; i < N; i++ {
		a[i] = N - 1 - a[i]
	}
	a = reverseOrderInt(a)
	calc(rd, RD)
	I++
	for i := 0; i < N; i++ {
		a[i] = N - 1 - a[i]
	}
	calc(ru, RU)
	x := 0
	for i := 0; i < N; i++ {
		if LU[i]+RU[N-1-i] != LIS-1 {
			continue
		}
		if LD[i]+RD[N-1-i] != LDS-1 {
			continue
		}
		tmp := lu[i] * ru[N-1-i] % mod * ld[i] % mod * rd[N-1-i] % mod
		x = (x + tmp) % mod
	}
	if x == sum {
		fmt.Println(LIS + LDS - 1)
	} else {
		fmt.Println(LIS + LDS)
	}
}

func getmax(i int) P {
	s := 0
	sum := 0
	for i > 0 {
		if s < bit[i] {
			s = bit[i]
			sum = bit2[i]
		} else if s == bit[i] {
			sum = (sum + bit2[i]) % mod
		}
		i = (i & (i - 1))
	}
	return P{s, sum}
}

func change(i, x, s int) {
	for i <= N {
		if bit[i] < x {
			bit[i] = x
			bit2[i] = s
		} else if bit[i] == x {
			bit2[i] = (bit2[i] + s) % mod
		}
		i += (i & -i)
	}
}

var I int
var sum int
var LIS, LDS int

func calc(dp, DP []int) {
	for i := 0; i < N+1; i++ {
		bit[i] = 0
		bit2[i] = 0
	}
	for i := 0; i < N+1; i++ {
		vc[i] = make([]int, 0)
	}
	M := 0
	for i := 0; i < N; i++ {
		p := getmax(a[i])
		x := p.x
		if x != 0 {
			dp[i] = p.y
		} else {
			dp[i] = 1
		}
		change(a[i]+1, x+1, dp[i])
		vc[x+1] = append(vc[x+1], i)
		M = max(M, x+1)
		DP[i] = x
	}
	if I <= 1 {
		al := 0
		for _, x := range vc[M] {
			al = (al + dp[x]) % mod
		}
		sum = sum * al % mod
	}
	if I == 0 {
		LIS = M
	}
	if I == 1 {
		LDS = M
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
