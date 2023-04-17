package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const ILL = 2167167167167167167
	const INF = 2100000000

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	sum := 0
	p := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &p[i])
		sum += p[i]
		p[i] = M - 2*p[i]
	}
	q := make([]int, 1)
	q[0] = -1
	Len := 1
	for i := 0; i < N; i++ {
		if p[i] == 0 {
			continue
		}
		if (!(q[Len-1] < 0) && (p[i] < 0)) || ((q[Len-1] < 0) && !(p[i] < 0)) {
			q = append(q, p[i])
			Len++
		} else {
			q[Len-1] += p[i]
		}
	}
	if Len%2 != 0 {
		resize(&q, Len-1)
		Len--
	}
	if K == 0 {
		fmt.Println(sum)
		return
	}
	if Len/2 <= K {
		for i := 0; i < Len/2; i++ {
			sum += q[i*2+1]
		}
		fmt.Println(sum)
		return
	}
	L := Len/2 + 1
	A := make([]int, L)
	B := make([]int, L)
	for i := 0; i < Len/2; i++ {
		A[i+1] = B[i] + q[i*2]
		B[i+1] = A[i+1] + q[i*2+1]
	}
	l := N * M
	r := 0

	var f func(int) pair
	f = func(val int) pair {
		ans := make([]pair, 2)
		ans[0] = pair{-ILL, -INF}
		ans[1] = pair{0, 0}
		for i := 1; i < L; i++ {
			tmp := ans[1]
			tmp.x += -A[i] - val
			tmp.y++
			ans[0] = maxPair(ans[0], tmp)
			tmp = ans[0]
			tmp.x += B[i]
			ans[1] = maxPair(ans[1], tmp)
		}
		return ans[1]
	}

	for l-r > 1 {
		med := (l + r) / 2
		if f(med).y < K {
			l = med
		} else {
			r = med
		}
	}
	tmp := f(r)
	fmt.Println(sum + tmp.x + K*r)
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

func maxPair(a, b pair) pair {
	if a.x > b.x {
		return a
	}
	return b
}
