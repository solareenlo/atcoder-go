package main

import (
	"fmt"
)

func main() {
	const INF = 1012345678

	var dpl, dpr [55][55][55]int

	var m int
	var s string
	fmt.Scan(&m, &s)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '+' || s[i] == '-' {
			dpl[i][i+1][0] = INF
		} else {
			dpl[i][i+1][0] = int(s[i] - 48)
		}
		dpl[i][i+1][1] = 0
		if s[i] == '+' || s[i] == '-' {
			dpr[i][i+1][0] = -INF
		} else {
			dpr[i][i+1][0] = int(s[i] - 48)
		}
		dpr[i][i+1][1] = 9
	}
	for i := 3; i <= n; i += 2 {
		for l := 0; l <= n-i; l++ {
			r := l + i
			for k := 0; k < i+1; k++ {
				dpl[l][r][k] = INF
				dpr[l][r][k] = -INF
			}
			var pc int
			if s[r-1] == '+' {
				pc = 0
			} else {
				pc = 1
			}
			var mc int
			if s[r-1] == '-' {
				mc = 0
			} else {
				mc = 1
			}
			for j := 1; j < i; j += 2 {
				for x := 0; x <= j; x++ {
					for y := 0; y <= i-j-1; y++ {
						dpl[l][r][x+y+pc] = min(dpl[l][r][x+y+pc], dpl[l][l+j][x]+dpl[l+j][r-1][y])
						dpl[l][r][x+y+mc] = min(dpl[l][r][x+y+mc], dpl[l][l+j][x]-dpr[l+j][r-1][y])
						dpr[l][r][x+y+pc] = max(dpr[l][r][x+y+pc], dpr[l][l+j][x]+dpr[l+j][r-1][y])
						dpr[l][r][x+y+mc] = max(dpr[l][r][x+y+mc], dpr[l][l+j][x]-dpl[l+j][r-1][y])
					}
				}
			}
		}
	}
	ret := dpr[0][n][0]
	for i := 0; i < m+1; i++ {
		ret = max(ret, dpr[0][n][i])
	}
	ok := (n%2 != 0 && ret >= -(INF/2))
	if ok {
		fmt.Println("OK")
		fmt.Println(ret)
	} else {
		fmt.Println("NG")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
