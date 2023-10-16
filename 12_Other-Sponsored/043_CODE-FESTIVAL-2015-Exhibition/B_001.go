package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)

	S := [2]int{}
	C := [2][100010]int{}
	N := 0
	D := [2][4][3]int{}
	L := [2][4]int{}

	fmt.Fscan(in, &S[0], &S[1], &N)

	for k := 0; k < 2; k++ {
		for j := 0; j < S[k]; j++ {
			C[k][j] = -1
		}
	}

	for N > 0 {
		x := [2]int{}
		c := 0
		fmt.Fscan(in, &x[0], &x[1], &c)
		x[0]--
		x[1]--
		c--

		if c >= 2 {
			c ^= 1
		}

		for k := 0; k < 2; k++ {
			var b int
			if (c & (1 << k)) > 0 {
				b = 1 ^ (x[1^k] & 1)
			} else {
				b = 0 ^ (x[1^k] & 1)
			}
			if C[k][x[k]] == -1 {
				C[k][x[k]] = b
			} else if C[k][x[k]] != b {
				fmt.Println("0")
				return
			}
		}
		N--
	}

	for k := 0; k < 2; k++ {
		D[k][0][2] = 1
		for i := 0; i < S[k]; i++ {
			tmp := [4][3]int{}

			for c := 0; c < 2; c++ {
				if C[k][i] == -1 || C[k][i] == c {
					for b := 0; b < 4; b++ {
						for l := 0; l < 3; l++ {
							if D[k][b][l] != 0 {
								p := 0
								if l == c {
									p = 1 << ((l + i) % 2)
								}
								nb := b | p
								nl := c
								tmp[nb][nl] = (tmp[nb][nl] + D[k][b][l]) % mod
							}
						}
					}
				}
			}

			for b := 0; b < 4; b++ {
				for l := 0; l < 3; l++ {
					D[k][b][l] = tmp[b][l]
				}
			}
		}

		for b := 0; b < 4; b++ {
			for l := 0; l < 3; l++ {
				L[k][b] = (L[k][b] + D[k][b][l]) % mod
			}
		}
	}

	ans := 0
	for b1 := 0; b1 < 4; b1++ {
		for b2 := 0; b2 < 4; b2++ {
			if (b1 & b2) == 0 {
				ans = (ans + L[0][b1]*L[1][b2]%mod) % mod
			}
		}
	}
	fmt.Printf("%d\n", ans)
}
