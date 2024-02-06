package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

var n, sg, gs int
var a, b, f [505][505]int

func swap_r(x, y int) {
	for i := 1; i <= n; i++ {
		a[x][i], a[y][i] = a[y][i], a[x][i]
	}
}

func swap_c(x, y int) {
	for i := 1; i <= n; i++ {
		a[i][x], a[i][y] = a[i][y], a[i][x]
	}
}

func add_r(x, y, z int) {
	for i := 1; i <= n; i++ {
		a[y][i] = (a[y][i] + a[x][i]*z) % MOD
	}
}

func add_c(x, y, z int) {
	for i := 1; i <= n; i++ {
		a[i][y] = (a[i][y] + a[i][x]*z) % MOD
	}
}

func Inver(x int) int {
	ans := 1
	y := MOD - 2
	for y > 0 {
		if (y & 1) != 0 {
			ans = ans * x % MOD
		}
		y >>= 1
		x = x * x % MOD
	}
	return ans
}

func Add(x, y, z int) {
	for i := 1; i <= n; i++ {
		a[y][i] = (a[y][i] + a[x][i]*z) % MOD
		b[y][i] = (b[y][i] + b[x][i]*z) % MOD
	}
}

func Gauss() {
	for i := 1; i <= n; i++ {
		for b[i][i] == 0 {
			for j := i + 1; j <= n; j++ {
				if b[i][j] != 0 {
					sg = -sg
					for k := 1; k <= n; k++ {
						a[k][i], a[k][j] = a[k][j], a[k][i]
						b[k][i], b[k][j] = b[k][j], b[k][i]
					}
					break
				}
			}
			if b[i][i] == 0 {
				gs++
				if gs > n {
					break
				}
				for j := 1; j <= n; j++ {
					b[i][j] = a[i][j]
					a[i][j] = 0
				}
				for j := 1; j < i; j++ {
					Add(j, i, -b[i][j])
				}
			}
		}
		if gs > n {
			break
		}
		ny := Inver(b[i][i])
		sg = sg * b[i][i] % MOD
		for j := 1; j <= n; j++ {
			a[i][j] = a[i][j] * ny % MOD
			b[i][j] = b[i][j] * ny % MOD
		}
		for j := 1; j <= n; j++ {
			if j != i && b[j][i] != 0 {
				Add(i, j, -b[j][i])
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			a[i][j] = -a[i][j]
		}
	}
}

func Hessenburg() {
	for i := 1; i < n; i++ {
		if a[i+1][i] == 0 {
			flag := 0
			for j := i + 2; j <= n; j++ {
				if a[j][i] != 0 {
					flag = 1
					swap_r(i+1, j)
					swap_c(i+1, j)
					break
				}
			}
			if flag == 0 {
				continue
			}
		}
		x := Inver(a[i+1][i])
		for j := i + 2; j <= n; j++ {
			if a[j][i] != 0 {
				xz := x * a[j][i] % MOD
				add_r(i+1, j, -xz)
				add_c(j, i+1, xz)
			}
		}
	}
}

func Det() {
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			f[i][j] = -a[i][i] * f[i-1][j] % MOD
		}
		for j := 1; j <= i; j++ {
			f[i][j] = (f[i][j] + f[i-1][j-1]) % MOD
		}
		S := -a[i][i-1]
		for j := i - 2; j >= 0; j-- {
			cs := S * a[j+1][i] % MOD
			for k := 0; k <= j; k++ {
				f[i][k] = (f[i][k] + f[j][k]*cs) % MOD
			}
			S = S * a[j+1][j] % MOD
		}
	}
}

var p [505]int

var out = bufio.NewWriter(os.Stdout)

func main() {
	in := bufio.NewReader(os.Stdin)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if p[i] > p[j] {
				b[i][i]++
				b[j][j]++
				b[i][j]--
				b[j][i]--
			} else {
				a[i][i]++
				a[j][j]++
				a[i][j]--
				a[j][i]--
			}
		}
	}
	n--
	gs = 0
	sg = 1
	Gauss()
	Hessenburg()
	Det()
	for i := 0; i <= n; i++ {
		aa := (sg*f[n][min(n+1, i+gs)]%MOD + MOD) % MOD
		fmt.Fprintf(out, "%d", aa)
		if i == n {
			fmt.Fprintf(out, "\n")
		} else {
			fmt.Fprintf(out, " ")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
