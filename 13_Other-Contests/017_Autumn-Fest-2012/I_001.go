package main

import "fmt"

const p = int(1e9 + 7)
const mod = int(1e9 + 9)
const o = 1010

var mat [3 * o][3 * o]int
var a, ma, aa int
var b, mb, ab int
var d, md, ad int

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)
	fmt.Scan(&a, &ma, &aa)
	fmt.Scan(&b, &mb, &ab)
	fmt.Scan(&d, &md, &ad)
	for i := 0; i < q; i++ {
		add(a-d+1, b)
		sub(a-d+1, b+1)
		add(a-d+2, b)
		sub(a-d+2, b+1)

		add(a+1, b+d)
		add(a+1, b+d+1)
		sub(a+2, b+d)
		sub(a+2, b+d+1)

		sub(a+d+1, b)
		sub(a+d+2, b)
		add(a+d+1, b+1)
		add(a+d+2, b+1)

		sub(a+1, b-d)
		sub(a+1, b-d+1)
		add(a+2, b-d)
		add(a+2, b-d+1)

		a = addMod(mul(a, ma, n), aa, n) + 1
		b = addMod(mul(b, mb, m), ab, m) + 1
		d = addMod(mul(d, md, max(m, n)), ad, max(m, n)) + 1
	}

	for i := 0; i < 3*o; i++ {
		for j := 0; j < 3*o; j++ {
			if i > 0 && j > 0 {
				mat[i][j] += mat[i-1][j-1]
				mat[i][j] %= mod
			}
		}
	}

	for i := 0; i < 3*o; i++ {
		for j := 3*o - 1; j >= 0; j-- {
			if i > 0 && j < 3*o-1 {
				mat[i][j] += mat[i-1][j+1]
				mat[i][j] %= mod
			}
		}
	}

	for i := 0; i < 3*o; i++ {
		for j := 0; j < 3*o; j++ {
			if i > 0 {
				mat[i][j] += mat[i-1][j]
				mat[i][j] %= mod
			}
		}
	}

	for i := 0; i < 3*o; i++ {
		for j := 0; j < 3*o; j++ {
			if j > 0 {
				mat[i][j] += mat[i][j-1]
				mat[i][j] %= mod
			}
		}
	}

	hash := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			hash = addMod(mul(hash, p, mod), get(i, j), mod)
		}
	}

	fmt.Println(hash)
}

func addMod(a, b, m int) int {
	return (a + b) % m
}

func mul(a, b, m int) int {
	return a * b % m
}

func add(x, y int) {
	mat[o+x][o+y]++
}

func sub(x, y int) {
	mat[o+x][o+y]--
}

func get(x, y int) int {
	return mat[o+x][o+y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
