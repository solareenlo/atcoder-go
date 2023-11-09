package main

import "fmt"

func main() {
	const P = 998244353
	const N = 61

	var f [N][11][11][11][2][2][2]int
	var pw [N]int
	for i := 0; i < N; i++ {
		pw[i] = 1 << i
	}

	var inc func(*int, int)
	inc = func(x *int, y int) {
		(*x) = ((*x) + y) % P
	}

	var n, a, b, c int
	fmt.Scan(&n, &a, &b, &c)
	f[N-1][0][0][0][0][0][0] = 1
	for i := N - 2; i >= 0; i-- {
		u := ((n & pw[i]) >> i)
		for x := 0; x < a; x++ {
			for y := 0; y < b; y++ {
				for z := 0; z < c; z++ {
					for p := 0; p < 2; p++ {
						for q := 0; q < 2; q++ {
							for r := 0; r < 2; r++ {
								v := f[i+1][x][y][z][p][q][r]
								if v == 0 {
									continue
								}
								if u != 0 {
									inc(&f[i][x][y][z][1][1][1], v)
								} else {
									inc(&f[i][x][y][z][p][q][r], v)
								}
								if u != 0 {
									inc(&f[i][(pw[i]+x)%a][(pw[i]+y)%b][z][p][q][1], v)
									inc(&f[i][x][(pw[i]+y)%b][(pw[i]+z)%c][1][q][r], v)
									inc(&f[i][(pw[i]+x)%a][y][(pw[i]+z)%c][p][1][r], v)
								} else {
									if p != 0 && q != 0 {
										inc(&f[i][(pw[i]+x)%a][(pw[i]+y)%b][z][p][q][r], v)
									}
									if q != 0 && r != 0 {
										inc(&f[i][x][(pw[i]+y)%b][(pw[i]+z)%c][p][q][r], v)
									}
									if p != 0 && r != 0 {
										inc(&f[i][(pw[i]+x)%a][y][(pw[i]+z)%c][p][q][r], v)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	ans := 0
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			for r := 0; r < 2; r++ {
				inc(&ans, f[0][0][0][0][p][q][r])
			}
		}
	}
	inc(&ans, P-1)
	inc(&ans, P-(n/(a*b/gcd(a, b)))%P)
	inc(&ans, P-(n/(a*c/gcd(a, c)))%P)
	inc(&ans, P-(n/(c*b/gcd(b, c)))%P)
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
