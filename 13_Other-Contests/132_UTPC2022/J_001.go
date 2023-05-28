package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	P := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &P[i])
	}

	ans := make([]pair, 0)

	var f func([]int) int
	f = func(P []int) int {
		res := 0
		for _, p := range P {
			res *= 10
			res += p
		}
		return res
	}

	var naive func() bool
	naive = func() bool {
		for i := 0; i < n; i++ {
			P[i]--
		}
		bs := make(map[int]pair)
		bef := make(map[int]int)
		que := make([]int, 0)
		que = append(que, f(P))
		bef[f(P)] = -1
		for len(que) != 0 {
			tmp := que[0]
			x := tmp
			que = que[1:]
			for i := n - 1; i >= 0; i-- {
				P[i] = tmp % 10
				tmp /= 10
			}
			for l := 0; l < n; l++ {
				for r := l + 4; r < n; r += 2 {
					Q := make([]int, len(P))
					copy(Q, P)
					d := r - l + 1
					tmp := P[l : r+1]
					V := make([]int, len(tmp))
					copy(V, tmp)
					sort.Ints(V)
					m := V[d/2]
					mid := -1
					for i := l; i < r+1; i++ {
						if P[i] == m {
							mid = i
						}
					}
					sort.Ints(Q[l:mid])
					sort.Ints(Q[mid+1 : r+1])
					h := f(Q)
					if _, ok := bef[h]; !ok {
						bef[h] = x
						bs[h] = pair{l + 1, r + 1}
						que = append(que, h)
					}
				}
			}
		}
		h := 0
		for i := 0; i < n; i++ {
			h = h*10 + i
		}
		if _, ok := bef[h]; !ok {
			return false
		}
		for bef[h] != -1 {
			ans = append(ans, bs[h])
			h = bef[h]
		}
		ans = reverseOrderPair(ans)
		return true
	}
	if n <= 9 {
		if naive() {
			fmt.Println(len(ans))
			for _, tmp := range ans {
				fmt.Println(tmp.x, tmp.y)
			}
		} else {
			fmt.Println(-1)
		}
		return
	}

	var query func(int, int)
	query = func(l, r int) {
		ans = append(ans, pair{l + 1, r + 1})
		d := r - l + 1
		tmp := P[l : r+1]
		V := make([]int, len(tmp))
		copy(V, tmp)
		sort.Ints(V)
		m := V[d/2]
		mid := -1
		for i := l; i < r+1; i++ {
			if P[i] == m {
				mid = i
			}
		}
		sort.Ints(P[l:mid])
		sort.Ints(P[mid+1 : r+1])
	}

	var issorted func(int) bool
	issorted = func(n int) bool {
		for i := 0; i < n-1; i++ {
			if P[i] > P[i+1] {
				return false
			}
		}
		return true
	}

	var insert func(int)
	insert = func(n int) {
		p := P[n-1]
		if p == n {
			return
		}
		if p == (n+1)/2 {
			l := 1 + (n & 1)
			query(l, n-1)
		} else if n%2 == 1 {
			query(0, n-1)
		} else {
			query(1, n-1)
		}
		if issorted(n) {
			return
		}
		p = n / 2
		query(p-1, p+3)
		if n%2 == 1 {
			query(0, n-1)
		} else {
			query(p-2, p+2)
			if issorted(n) {
				return
			}
			query(0, n-2)
		}
	}

	var solveodd func(int)
	solveodd = func(n int) {
		query(0, n-1)
		k := n / 2
		m := n/2 + 1
		if P[n-1] == m {
			insert(n)
			return
		}
		if P[n-2] > m {
			query(0, n-3)
			query(0, n-1)
			if issorted(n - 1) {
				return
			}
			query(k-1, k+3)
			query(0, n-1)
			return
		} else {
			pi := -1
			for i := 0; i < n; i++ {
				if P[i] == m {
					pi = i
					break
				}
			}
			if pi >= 4 {
				query(pi-4, pi)
			} else {
				query(pi-3, pi+1)
			}
			query(0, n-1)
			query(0, n-3)
			query(0, n-1)
			if issorted(n - 1) {
				return
			}
			query(k-1, k+3)
			query(0, n-1)
			return
		}
	}

	if n%2 == 0 {
		m := P[n-1]
		for i := 0; i < n-1; i++ {
			if P[i] > m {
				P[i]--
			}
		}
		solveodd(n - 1)
		for i := 0; i < n-1; i++ {
			if P[i] >= m {
				P[i]++
			}
		}
		insert(n)
	} else {
		solveodd(n)
	}

	fmt.Println(len(ans))
	for _, tmp := range ans {
		fmt.Println(tmp.x, tmp.y)
	}
}

type pair struct {
	x, y int
}

func reverseOrderPair(a []pair) []pair {
	n := len(a)
	res := make([]pair, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
