package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

const mod1 = 999999797
const INF = 1 << 30

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const MAX = 100005

	var nex [26][MAX]int

	var S string
	fmt.Fscan(in, &S)
	N := len(S)
	for i := range nex {
		for j := range nex[i] {
			nex[i][j] = -1
		}
	}
	for i := 0; i < N; i++ {
		nex[S[i]-'a'][i] = i
	}

	rol := make([]Rollinghash, 26)

	aaa := r.Int() % 47382190
	bbb := r.Int() % 47382190

	for t := 0; t < 26; t++ {
		T := strings.Split(strings.Repeat("A", N), "")
		for i := 0; i < N; i++ {
			if S[i] == byte('a'+t) {
				T[i] = "B"
			}
		}
		rol[t].Make(strings.Join(T, ""), aaa, bbb)
	}

	for t := 0; t < 26; t++ {
		for i := N - 1; i >= 0; i-- {
			if nex[t][i] == -1 {
				nex[t][i] = nex[t][i+1]
			}
		}
	}

	ma := 0

	for s := 1; s < N; s++ {
		A := make([]pair, 26)
		B := make([]pair, 26)
		for i := 0; i < 26; i++ {
			A[i] = pair{nex[i][ma], i}
			B[i] = pair{nex[i][s], i}
			if A[i].x == -1 {
				A[i].x = INF
			}
			if B[i].x == -1 {
				B[i].x = INF
			}
		}
		sortPair(A)
		sortPair(B)
		dicA := make([]int, 26)
		dicB := make([]int, 26)
		for i := 0; i < 26; i++ {
			dicA[A[i].y] = i
			dicB[B[i].y] = i
		}
		Len := N - s
		for i := 0; i < 26; i++ {
			a := A[i].y
			b := B[i].y
			if A[i].x >= ma+Len && B[i].x >= s+Len {
				break
			}
			ha1 := rol[a].Hash(ma, ma+Len)
			ha2 := rol[b].Hash(s, s+Len)
			if ha1 == ha2 {
				continue
			} else {
				l := 0
				r := Len
				for r-l > 1 {
					m := (l + r) / 2
					aa := rol[a].Hash(ma, ma+m)
					bb := rol[b].Hash(s, s+m)
					if aa == bb {
						l = m
					} else {
						r = m
					}
				}
				Len = min(Len, l)
			}
		}

		if Len != N-s {
			x := dicA[S[ma+Len]-'a']
			y := dicB[S[s+Len]-'a']
			if x < y {
			} else {
				ma = s
			}
		}
	}

	A := make([]pair, 26)
	for i := 0; i < 26; i++ {
		A[i] = pair{nex[i][ma], i}
		if A[i].x == -1 {
			A[i].x = INF
		}
	}

	sortPair(A)

	dicA := make([]int, 26)
	for i := 0; i < 26; i++ {
		dicA[A[i].y] = i
	}

	for s := ma; s < N; s++ {
		fmt.Fprint(out, string('a'+(25-dicA[S[s]-'a'])))
	}
	fmt.Fprintln(out)
}

type Rollinghash struct {
	S               string
	n, base1, base2 int
	h1, ru1         []int
}

func (hash *Rollinghash) Make(T string, ba1, ba2 int) {
	hash.S = T
	hash.n = len(hash.S)
	hash.h1 = make([]int, hash.n+1)
	hash.ru1 = make([]int, hash.n+1)
	hash.base1 = ba1
	hash.ru1[0] = 1
	for i := 1; i <= hash.n; i++ {
		hash.h1[i] = hash.h1[i-1]*hash.base1 + int(hash.S[i-1]-'A')
		hash.h1[i] %= mod1
		hash.ru1[i] = hash.ru1[i-1] * hash.base1 % mod1
	}
}

func (hash Rollinghash) Hash(l, r int) int {
	return (hash.h1[r] - hash.h1[l]*hash.ru1[r-l]%mod1 + mod1) % mod1
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
