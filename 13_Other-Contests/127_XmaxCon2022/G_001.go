package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var arr []int
var vc [1000007][]int
var vis [1000007]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n, s int
		fmt.Fscan(in, &n, &s)
		arr = Generate(n, s)
		for i := 1; i < n+1; i++ {
			vc[i] = make([]int, 0)
		}
		for pos := 0; pos < n; pos++ {
			ban := make([]int, 0)
			if Calc(pos) == 0 {
				ban = append(ban, pos)
			}
			for i := 0; i < 4; i++ {
				vc[abs(arr[pos*4+i])] = append(vc[abs(arr[pos*4+i])], pos)
			}
			for len(ban) != 0 {
				u := ban[0]
				ban = ban[1:]
				rng := r.Int()%4 + u*4
				ap := abs(arr[rng])
				if vis[ap] {
					vis[ap] = false
				} else {
					vis[ap] = true
				}
				for _, u := range vc[ap] {
					if Calc(u) == 0 {
						ban = append(ban, u)
					}
				}
			}
		}
		for i := 1; i < n+1; i++ {
			if vis[i] {
				fmt.Fprint(out, "+")
			} else {
				fmt.Fprint(out, "-")
			}
		}
		fmt.Fprintln(out)
	}
}

func Generate(N, S int) []int {
	s := S
	A := make([]int, 4*N)
	for i := 0; i < 4*N; i++ {
		A[i] = i/4 + 1
	}
	for i := 0; i < 4*N; i++ {
		s = (s * 2022) % 998244353
		if s%2 != 0 {
			A[i] = -A[i]
		}
	}
	for i := 0; i < 4*N; i++ {
		s = (s * 2022) % 998244353
		j := s % (i + 1)
		t := A[j]
		A[j] = A[i]
		A[i] = t
	}
	return A
}

func Calc(u int) int {
	u *= 4
	cnt := 0
	for i := 0; i < 4; i++ {
		if vis[abs(arr[u])] == (arr[u] > 0) {
			cnt++
		}
		u++
	}
	return cnt
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
