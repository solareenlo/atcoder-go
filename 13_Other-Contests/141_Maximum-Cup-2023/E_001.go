package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var N int
	fmt.Fscan(in, &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
	}

	lmin := make([]int, N)
	lmax := make([]int, N)
	rmin := make([]int, N)
	rmax := make([]int, N)
	st1 := make([]int, 0)
	st2 := make([]int, 0)
	st3 := make([]int, 0)
	st4 := make([]int, 0)
	for i := 0; i < N; i++ {
		for len(st1) != 0 && P[st1[len(st1)-1]] > P[i] {
			st1 = st1[:len(st1)-1]
		}
		if len(st1) != 0 {
			lmin[i] = st1[len(st1)-1]
		} else {
			lmin[i] = -1
		}
		st1 = append(st1, i)
		for len(st2) != 0 && P[st2[len(st2)-1]] < P[i] {
			st2 = st2[:len(st2)-1]
		}
		if len(st2) != 0 {
			lmax[i] = st2[len(st2)-1]
		} else {
			lmax[i] = -1
		}
		st2 = append(st2, i)
	}
	for i := N - 1; i >= 0; i-- {
		for len(st3) != 0 && P[st3[len(st3)-1]] > P[i] {
			st3 = st3[:len(st3)-1]
		}
		if len(st3) != 0 {
			rmin[i] = st3[len(st3)-1]
		} else {
			rmin[i] = N
		}
		st3 = append(st3, i)
		for len(st4) != 0 && P[st4[len(st4)-1]] < P[i] {
			st4 = st4[:len(st4)-1]
		}
		if len(st4) != 0 {
			rmax[i] = st4[len(st4)-1]
		} else {
			rmax[i] = N
		}
		st4 = append(st4, i)
	}

	dp := make([][2]int, N)
	subdp := make([][2]int, N+2)
	dpsum := make([]int, 2)
	for i := 0; i < N; i++ {
		if lmin[i] == -1 {
			dp[i][0] = 1
		}
		dp[i][0] = (dp[i][0] + (dpsum[1]-subdp[lmin[i]+1][0]+MOD)%MOD) % MOD
		dp[i][1] = (dp[i][1] + (dpsum[0]-subdp[lmax[i]+1][1]+MOD)%MOD) % MOD
		subdp[rmax[i]+1][0] = (subdp[rmax[i]+1][0] + dp[i][1]) % MOD
		subdp[rmin[i]+1][1] = (subdp[rmin[i]+1][1] + dp[i][0]) % MOD
		subdp[i+1][0] = (subdp[i+1][0] + subdp[i][0]) % MOD
		subdp[i+1][1] = (subdp[i+1][1] + subdp[i][1]) % MOD
		dpsum[0] = (dpsum[0] + dp[i][0]) % MOD
		dpsum[1] = (dpsum[1] + dp[i][1]) % MOD
	}

	ans := 0
	for i := 0; i < N; i++ {
		if rmin[i] == N {
			ans = (ans + dp[i][0]) % MOD
		}
		if rmax[i] == N {
			ans = (ans + dp[i][1]) % MOD
		}
	}
	fmt.Println(ans)
}
