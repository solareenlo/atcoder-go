package main

import (
	"fmt"
)

func check(a, b string) bool {
	for i := 0; i <= len(a); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func main() {
	var a, b string
	var c, d int
	fmt.Scan(&a, &b, &c, &d)
	Len := len(a)
	if check(a, b) {
		fmt.Println(0)
		return
	}

	m := make([]int, 85)
	m[0] = 1
	for i := 1; i <= 80; i++ {
		m[i] = m[i-1] * 10 % c
	}

	n := make([]int, 85)
	for i := 0; i < Len/2; i++ {
		n[i] = (m[i] + m[Len-i-1]) % c
	}

	var dp [85][4][55][725]int
	for i := a[0] - 48; i <= b[0]-48; i++ {
		u := 0
		if i == a[0]-48 {
			u++
		}
		if i == b[0]-48 {
			u += 2
		}
		dp[1][u][(n[0]*int(i))%c][i*2]++
	}

	p := 10000
	for i := 1; i < Len/2; i++ {
		for j := 0; j <= 3; j++ {
			for k := 0; k < c; k++ {
				for l := 0; l <= d; l++ {
					if dp[i][j][k][l] != 0 {
						for o := 0; o <= 9; o++ {
							if (j%2 != 0 && o < int(a[i]-48)) || (j/2 != 0 && o > int(b[i]-48)) {
								continue
							}
							z := j
							if o > int(a[i]-48) {
								z -= z % 2
							}
							if o < int(b[i]-48) && z >= 2 {
								z -= 2
							}
							dp[i+1][z][(k+n[i]*o)%c][(l + o*2)] = (dp[i+1][z][(k+n[i]*o)%c][(l+o*2)] + dp[i][j][k][l]) % p
						}
					}
				}
			}
		}
	}

	sum := 0
	for i := 0; i <= 3; i++ {
		sum = (sum + dp[Len/2][i][0][d]) % p
	}
	fmt.Println(sum)
}
