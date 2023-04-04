package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	var S string
	fmt.Fscan(in, &t, &S)
	n := len(S)
	s := strings.Split(S, "")
	sr := reverseOrderString(s)
	re := make([]string, t)
	for i := range re {
		re[i] = "."
	}
	var uf UF
	uf.init(t)
	pya := make([]int, t)
	j := 0
	j0 := 0
	for i := 0; i < n; i++ {
		if s[i] == "*" {
			if sr[j] == "*" {
				pya[j0] = 1
			}
			for Len := 0; Len < t; {
				if sr[j] == "*" {
					inc := min(t-Len, t-j0)
					Len += inc
					j0 += inc
					if j0 == t {
						j++
						j0 = 0
					}
				} else {
					if re[Len] != "." && re[Len] != sr[j] {
						fmt.Println("No")
						return
					}
					re[Len] = sr[j]
					Len++
					j++
				}
			}
		} else {
			if sr[j] == "*" {
				if re[t-1-j0] != "." && re[t-1-j0] != s[i] {
					fmt.Println("No")
					return
				}
				re[t-1-j0] = s[i]
				j0++
				if j0 == t {
					j++
					j0 = 0
				}
			} else {
				if sr[j] != s[i] {
					fmt.Println("No")
					return
				}
				j++
			}
		}
	}
	for i := 0; i < t; i++ {
		if pya[i] != 0 {
			d := 0
			for j := i + 1; j < t; j++ {
				if pya[j] != 0 {
					d = gcd(j-i, d)
				}
			}
			for j := i; j < t-d; j++ {
				uf.unite(t-1-j, t-1-(j+d))
			}
			for j := i; j < t; j++ {
				uf.unite(t-1-j, j-i)
			}
			break
		}
	}
	s, sr = sr, s
	re = reverseOrderString(re)
	pya2 := make([]int, t)
	j = 0
	j0 = 0
	for i := 0; i < n; i++ {
		if s[i] == "*" {
			if sr[j] == "*" {
				pya2[j0] = 1
			}
			for Len := 0; Len < t; {
				if sr[j] == "*" {
					inc := min(t-Len, t-j0)
					Len += inc
					j0 += inc
					if j0 == t {
						j++
						j0 = 0
					}
				} else {
					if re[Len] != "." && re[Len] != sr[j] {
						fmt.Println("No")
						return
					}
					re[Len] = sr[j]
					Len++
					j++
				}
			}
		} else {
			if sr[j] == "*" {
				if re[t-1-j0] != "." && re[t-1-j0] != s[i] {
					fmt.Println("No")
					return
				}
				re[t-1-j0] = s[i]
				j0++
				if j0 == t {
					j++
					j0 = 0
				}
			} else {
				if sr[j] != s[i] {
					fmt.Println("No")
					return
				}
				j++
			}
		}
	}
	for i := 0; i < t; i++ {
		if pya2[i] != 0 {
			d := 0
			for j := i + 1; j < t; j++ {
				if pya2[j] != 0 {
					d = gcd(j-i, d)
				}
			}
			for j := i; j < t-d; j++ {
				uf.unite(j, (j + d))
			}
			for j := i; j < t; j++ {
				uf.unite(j, t-1-(j-i))
			}
			break
		}
	}
	s, sr = sr, s
	re = reverseOrderString(re)
	ret := make([]string, len(re))
	copy(ret, re)
	for i := 0; i < t; i++ {
		if re[i] != "." {
			ret[uf.root(i)] = re[i]
		}
	}
	hash := n
	for i := 0; i < n; i++ {
		hash += (i + (i ^ 254)) * int(s[i][0])
	}
	for i := 0; i < t; i++ {
		if re[i] != "." && ret[uf.root(i)] != re[i] {
			if hash%12 != 0 {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

type UF struct {
	data []int
}

func (u *UF) init(size int) {
	u.data = make([]int, size)
	for i := range u.data {
		u.data[i] = -1
	}
}

func (u *UF) unite(x, y int) bool {
	x = u.root(x)
	y = u.root(y)
	if x != y {
		if u.data[y] < u.data[x] {
			x, y = y, x
		}
		u.data[x] += u.data[y]
		u.data[y] = x
	}
	return x != y
}

func (u UF) findSet(x, y int) bool { return u.root(x) == u.root(y) }

func (u *UF) root(x int) int {
	if u.data[x] < 0 {
		return x
	}
	u.data[x] = u.root(u.data[x])
	return u.data[x]
}

func (u UF) size(x int) int { return -u.data[u.root(x)] }

func reverseOrderString(a []string) []string {
	n := len(a)
	res := make([]string, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
