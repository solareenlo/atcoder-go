package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 0x3f3f3f3f

func main() {
	in := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(in, &H, &W)
	c := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &c[i])
	}
	roads := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if c[i][j] == '.' {
				roads++
			}
		}
	}
	pow3 := make([]int, W+1)
	pow3[0] = 1
	for i := 1; i <= W; i++ {
		pow3[i] = pow3[i-1] * 3
	}
	pow3W := pow3[W]
	pow3W1 := pow3[W-1]

	var multTable [256][256]int
	const IrreduciblePoly = (1 << 8) | (1 << 4) | (1 << 3) | (1 << 1) | 1
	for a := 0; a < 256; a++ {
		for b := 0; b < 256; b++ {
			multTable[a][b] = multBitPolynomial(a, b, IrreduciblePoly, 1<<7)
		}
	}

	minRoute := INF
	maxSize := H * W
	dp := make([][]int, 0)
	ndp := make([][]int, pow3W)
	for i := range ndp {
		ndp[i] = make([]int, maxSize+1)
	}
	ndp[0][0] = 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			cc := c[i][j]
			upc := byte('#')
			if i != 0 {
				upc = c[i-1][j]
			}
			leftc := byte('#')
			if j != 0 {
				leftc = c[i][j-1]
			}
			w := xor128() & 0xff
			for w == 0 {
				w = xor128() & 0xff
			}
			dp, ndp = ndp, dp
			ndp = make([][]int, pow3W)
			for k := range ndp {
				ndp[k] = make([]int, maxSize+1)
			}
			for k := 0; k < pow3W; k++ {
				for s := 0; s < maxSize; s++ {
					x := dp[k][s]
					if x == 0 {
						continue
					}
					tmpUP := 0
					if upc == 'H' {
						tmpUP = 3
					}
					up := tmpUP + (k / pow3W1 % 3)
					tmpLEFT := 0
					if leftc == 'H' {
						tmpLEFT = 3
					}
					tmpJ := 0
					if j != 0 {
						tmpJ = k % 3
					}
					left := tmpLEFT + tmpJ
					nk := k * 3 % pow3W
					if cc == '#' {
						if up != 4 {
							ndp[nk+0][s] ^= x
						}
					} else if cc == 'H' {
						if up != 4 {
							if up == 1 || up == 2 || left == 1 || left == 2 {
								ndp[nk+2][s] ^= x
							} else {
								ndp[nk+1][s] ^= x
							}
						}
					} else if cc == '.' {
						if up != 4 {
							ndp[nk+0][s] ^= x
						}
						tmp := nk
						if left == 4 {
							tmp = nk + 3
						}
						nk2 := tmp
						y := multTable[w][x]
						if up != 2 && left != 2 {
							ndp[nk2+1][s+1] ^= y
						}
						if up != 1 && left != 1 && s != 0 {
							ndp[nk2+2][s+1] ^= y
						}
					}
				}
			}
		}
	}
	curMinRoute := INF
	cnt := make([]int, maxSize+1)
	for k := 0; k < pow3W; k++ {
		ok := true
		for j := 0; j < W; j++ {
			cc := c[H-1][W-1-j]
			tmpH := 0
			if cc == 'H' {
				tmpH = 3
			}
			t := tmpH + (k / pow3[j] % 3)
			if ok && t != 4 {
				ok = true
			} else {
				ok = false
			}
		}
		if !ok {
			continue
		}
		for s := 0; s < maxSize; s++ {
			cnt[s] ^= ndp[k][s]
		}
	}
	for s := 0; s < maxSize; s++ {
		if cnt[s] != 0 {
			curMinRoute = min(curMinRoute, s)
		}
	}
	minRoute = min(minRoute, curMinRoute)
	if minRoute == INF {
		fmt.Println(-1)
	} else {
		ans := roads - minRoute
		fmt.Println(ans)
	}
}

func multBitPolynomial(a, b, c, h int) int {
	t := 0
	for b != 0 {
		if (b & 1) != 0 {
			t ^= a
		}
		b >>= 1
		s := a & h
		a <<= 1
		if s != 0 {
			a ^= c
		}
	}
	return t
}

var x int = 123456789
var y int = 362436069
var z int = 521288629
var w int = 88675123

func xor128() int {
	t := x ^ (x << 11)
	x = y
	y = z
	z = w
	w = w ^ (w >> 19) ^ (t ^ (t >> 8))
	return w
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
