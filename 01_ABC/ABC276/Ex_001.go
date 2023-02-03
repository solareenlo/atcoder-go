package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

const N = 2005

var id [N][N]int
var mm int
var f [N]*big.Int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	var a, b, c, d, e [N]int
	var s [N][N]int
	for i := range f {
		f[i] = new(big.Int)
	}
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i], &d[i], &e[i])
		if e[i] != 0 {
			s[a[i]][c[i]]++
			s[b[i]+1][c[i]]--
			s[a[i]][d[i]+1]--
			s[b[i]+1][d[i]+1]++
			add(a[i]-1, c[i]-1, i)
			add(a[i]-1, d[i], i)
			add(b[i], c[i]-1, i)
			add(b[i], d[i], i)
		}
	}
	for i := 1; i <= q; i++ {
		if e[i] != 0 {
			if e[i] == 2 {
				f[i].SetBit(f[i], mm+1, 1)
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			s[i][j] += s[i-1][j] + s[i][j-1] - s[i-1][j-1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i][j] > 0 {
				s[i][j] = 1
			} else {
				s[i][j] = 0
			}
		}
	}
	idx := 1
	var pos [N]int
	for i := 1; i <= mm; i++ {
		for j := idx; j <= q; j++ {
			if f[j].Bit(i) != 0 {
				f[idx], f[j] = f[j], f[idx]
				break
			}
		}
		if f[idx].Bit(i) == 0 {
			continue
		}
		for j := 1; j <= q; j++ {
			if j != idx && f[j].Bit(i) != 0 {
				f[j].Xor(f[j], f[idx])
			}
		}
		pos[idx] = i
		idx++
	}
	for i := idx; i <= q; i++ {
		if f[i].Bit(mm+1) != 0 {
			fmt.Fprintln(out, "No")
			return
		}
	}
	var ans [N * 4]int
	for i := 1; i < idx; i++ {
		ans[pos[i]] = int(f[i].Bit(mm + 1))
	}
	var x [N][N]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			x[i][j] = ans[id[i][j]]
		}
	}
	for i := n; i > 0; i-- {
		for j := n; j > 0; j-- {
			x[i][j] ^= x[i-1][j] ^ x[i][j-1] ^ x[i-1][j-1]
			if s[i][j] != 0 {
				x[i][j]++
			} else {
				x[i][j] = 0
			}
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			s[i][j] = 1 - s[i][j]
			s[i][j] += s[i-1][j] + s[i][j-1] - s[i-1][j-1]
		}
	}
	for i := 1; i <= q; i++ {
		if e[i] == 0 && (s[b[i]][d[i]]-s[a[i]-1][d[i]]-s[b[i]][c[i]-1]+s[a[i]-1][c[i]-1]) == 0 {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fprintf(out, "%d ", x[i][j])
		}
		fmt.Fprintln(out)
	}
}

func add(x, y, i int) {
	if x != 0 && y != 0 {
		if id[x][y] == 0 {
			mm++
			id[x][y] = mm
		}
		f[i].SetBit(f[i], id[x][y], 1)
	}
}
