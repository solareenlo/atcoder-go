package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	rot := 0
	A := make([]int, 0)
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		if x == 90 {
			A = append(A, 1)
		} else {
			A = append(A, 0)
		}
		if A[i] == 1 {
			rot++
		} else {
			rot--
		}
	}
	if rot != 4 {
		fmt.Fprintln(out, -1)
		return
	}

	ret := solve(A)
	for _, r := range ret {
		fmt.Fprintln(out, r.x, r.y)
	}
}

type pair struct {
	x, y int
}

func solve(A []int) []pair {
	if len(A) == 4 {
		return []pair{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	}
	i := 0
	N := len(A)

	for i = 0; i < N; i++ {
		if A[i] == 1 && A[(i+1)%N] == 0 {
			break
		}
	}
	B := make([]int, 0)
	B = append(B, A[(i+N-1)%N])
	for j := 0; j < N-3; j++ {
		B = append(B, A[(i+2+j)%N])
	}
	P := solve(B)
	R := make([]pair, 0)
	for P[0].x >= P[1].x {
		for p := range P {
			P[p].x, P[p].y = P[p].y, P[p].x
			P[p].x = -P[p].x
		}
	}

	for p := range P {
		P[p].x *= 2
		P[p].y *= 2
	}
	P = insertPair(&P, 1, pair{P[0].x + 1, P[0].y})
	P = insertPair(&P, 2, pair{P[0].x + 1, P[0].y + 1})
	P[3].y++
	Xs := make([]int, 0)
	Ys := make([]int, 0)
	for _, p := range P {
		Xs = append(Xs, p.x)
		Ys = append(Ys, p.y)
	}
	sort.Ints(Xs)
	sort.Ints(Ys)
	Xs = unique(Xs)
	Ys = unique(Ys)

	for j := 0; j < N; j++ {
		x := lowerBound(Xs, P[(j-(i-1)+N)%N].x)
		y := lowerBound(Ys, P[(j-(i-1)+N)%N].y)
		R = append(R, pair{x, y})
	}

	return R
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func insertPair(a *[]pair, index int, value pair) []pair {
	n := len(*a)
	if index < 0 {
		index = (index%n + n) % n
	}
	switch {
	case index == n:
		return append(*a, value)

	case index < n:
		*a = append((*a)[:index+1], (*a)[index:]...)
		(*a)[index] = value
		return *a

	case index < cap(*a):
		*a = (*a)[:index+1]
		for i := n; i < index; i++ {
			(*a)[i] = pair{0, 0}
		}
		(*a)[index] = value
		return *a

	default:
		b := make([]pair, index+1)
		if n > 0 {
			copy(b, *a)
		}
		b[index] = value
		return b
	}
}
