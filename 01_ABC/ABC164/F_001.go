package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 500

var (
	s   = make([][]int, 2)
	u   = make([][]uint64, 2)
	n   int
	val = [2][N]int{}
	d   = [N][N]int{}
)

func flip() {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			d[i][j], d[j][i] = d[j][i], d[i][j]
		}
	}
}

func solve() bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d[i][j] = -1
		}
	}
	for k := 0; k < 2; k++ {
		for i := 0; i < n; i++ {
			x := val[k][i]
			if s[k][i] != x {
				for j := 0; j < n; j++ {
					if d[i][j] == 0 && x == 1 {
						return false
					} else if d[i][j] == 1 && x == 0 {
						return false
					}
					d[i][j] = x
				}
			}
		}
		flip()
	}
	for l := 0; l < 2; l++ {
		for k := 0; k < 2; k++ {
			for i := 0; i < n; i++ {
				x := val[k][i]
				if s[k][i] == x {
					p := make([]int, 0)
					ok := false
					for j := 0; j < n; j++ {
						if d[i][j] == -1 {
							p = append(p, j)
						}
						if d[i][j] == x {
							ok = true
						}
					}
					if ok {
						continue
					}
					if len(p) == 0 {
						return false
					}
					if len(p) == 1 {
						d[i][p[0]] = x
					}
				}
			}
			flip()
		}
	}

	is := make([]int, 0)
	js := make([]int, 0)
	for i := 0; i < n; i++ {
		filled := true
		for j := 0; j < n; j++ {
			if d[i][j] == -1 {
				filled = false
			}
		}
		if !filled {
			is = append(is, i)
		}
	}
	for j := 0; j < n; j++ {
		filled := true
		for i := 0; i < n; i++ {
			if d[i][j] == -1 {
				filled = false
			}
		}
		if !filled {
			js = append(js, j)
		}
	}
	for i := 0; i < len(is); i++ {
		for j := 0; j < len(js); j++ {
			d[is[i]][js[j]] = (i + j) % 2
		}
	}
	return true
}

var ans = [N][N]uint64{}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 0; i < 2; i++ {
		s[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &s[i][j])
		}
	}
	for i := 0; i < 2; i++ {
		u[i] = make([]uint64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &u[i][j])
		}
	}

	for bit := 0; bit < 64; bit++ {
		for i := 0; i < 2; i++ {
			for j := 0; j < n; j++ {
				val[i][j] = int(u[i][j]>>bit) & 1
			}
		}
		if !solve() {
			fmt.Fprintln(out, -1)
			return
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ans[i][j] |= uint64(d[i][j] << bit)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(out, ans[i][j])
			if j == n-1 {
				fmt.Fprint(out, "\n")
			} else {
				fmt.Fprint(out, " ")
			}
		}
	}
}
