package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M int
var A [100][7]int
var mem [101][7]map[int]int

func dp(n, m, cp int) int {
	if _, ok := mem[n][m][cp]; ok {
		return mem[n][m][cp]
	}

	ret := 0
	nn := n + (m+1)/M
	nm := (m + 1) % M

	mx := 0
	tmp := cp
	var arr [8]int
	for m := 0; m < M; m++ {
		arr[m] = tmp % 10
		mx = max(mx, arr[m])
		if arr[m] > 1 {
			ret = 1e9
		}
		tmp /= 10
	}
	if n == N {
		return ret
	}

	chk := false
	for m := 1; m < M; m++ {
		if arr[0] == arr[m] {
			chk = true
		}
	}
	if chk || arr[0] == 0 {
		ret = min(ret, dp(nn, nm, cp/10))
	}

	if arr[0] == 0 {
		if m == 0 || arr[M-1] == 0 {
			arr[M] = mx + 1
		} else {
			arr[M] = arr[M-1]
		}
	} else {
		if m == 0 || arr[M-1] == 0 {
			arr[M] = arr[0]
		} else {
			for m := 1; m < M; m++ {
				if arr[m] == arr[M-1] {
					arr[m] = arr[0]
				}
			}
			arr[M] = arr[0]
		}
	}

	var idx [8]int
	cnt := 0
	for m := 1; m <= M; m++ {
		if arr[m] != 0 && idx[arr[m]] == 0 {
			cnt++
			idx[arr[m]] = cnt
		}
	}

	ncp := 0
	x := 1
	for m := 1; m <= M; m++ {
		ncp += x * idx[arr[m]]
		x *= 10
	}
	ret = min(ret, A[n][m]+dp(nn, nm, ncp))

	mem[n][m][cp] = ret
	return mem[n][m][cp]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = make(map[int]int)
		}
	}

	fmt.Fscan(in, &N, &M)
	for n := 0; n < N; n++ {
		var ch string
		fmt.Fscan(in, &ch)
		for m := 0; m < M; m++ {
			if ch[m] == '#' {
				A[n][m] = -1000
			} else {
				A[n][m] = 1
			}
		}
	}
	fmt.Println((1000 + dp(0, 0, 0)%1000) % 1000)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
