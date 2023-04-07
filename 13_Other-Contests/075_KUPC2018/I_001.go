package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var N, M, e int
var LRS, TLR [][3]int
var is_left_end []bool
var right_end, Bit []int
var event [][5]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &N, &M)
	N += 10
	LRS = make([][3]int, M)
	is_left_end = make([]bool, N)
	right_end = make([]int, N)
	Bit = make([]int, N)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &LRS[i][0], &LRS[i][1], &LRS[i][2])
		LRS[i][0] += 5
		LRS[i][1] += 5
	}
	var Q int
	fmt.Fscan(in, &Q)
	TLR = make([][3]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &TLR[i][0], &TLR[i][1], &TLR[i][2])
		TLR[i][1] += 5
		TLR[i][2] += 5
	}
	event = make([][5]int, 1<<20)
	add_interval(-1, 1, 2)
	add_interval(-1, N-2, N-1)

	for q := 0; q < Q; q++ {
		t := TLR[q][0]
		L := TLR[q][1]
		R := TLR[q][2] + 1
		for {
			k := get_sum(L - 1)
			a := find_kth_element(k + 1)
			b := right_end[a]
			if L <= a && b <= R {
				rm_interval(q, a, b)
			} else {
				break
			}
		}
		k := get_sum(L)
		a := find_kth_element(k)
		b := right_end[a]
		c := find_kth_element(k + 1)
		d := right_end[c]
		if t == 0 {
			if a <= L && R <= b {
				rm_interval(q, a, b)
				add_interval(q, a, L)
				add_interval(q, R, b)
			} else {
				if b > L {
					rm_interval(q, a, b)
					add_interval(q, a, L)
				}
				if c < R {
					rm_interval(q, c, d)
					add_interval(q, R, d)
				}
			}
		}
		if t == 1 {
			if a <= L && R <= b {
			} else {
				if b > L {
					rm_interval(q, a, b)
					L = a
				}
				if c < R {
					rm_interval(q, c, d)
					R = d
				}
				add_interval(q, L, R)
			}
		}
	}

	resize(&event, e)
	sort.Slice(event, func(i, j int) bool {
		return event[i][1] < event[j][1]
	})

	sort.Slice(LRS, func(i, j int) bool {
		return LRS[i][0] < LRS[j][0]
	})

	ev_idx := make([]int, N+1)
	for i := 0; i <= N; i++ {
		ev_idx[i] = lowerBoundEvent(event, i)
	}

	LRS_idx := make([]int, N+1)
	for i := 0; i <= N; i++ {
		LRS_idx[i] = lowerBoundLRS(LRS, i)
	}

	diff := make([]int, Q)
	for i := range Bit {
		Bit[i] = 0
	}
	for x := 1; x < N; x++ {
		for e := ev_idx[x]; e < ev_idx[x+1]; e++ {
			q, l, r := event[e][0], event[e][2], event[e][3]
			coef := event[e][4]
			diff[q] += coef * (get_sum(r) - get_sum(l))
		}
		for i := LRS_idx[x]; i < LRS_idx[x+1]; i++ {
			r, s := LRS[i][1]+1, LRS[i][2]
			Add(r, s)
		}
	}

	ans := make([]int, Q+1)
	for i := 0; i < Q; i++ {
		ans[i+1] = ans[i] + diff[i]
	}
	for i := 1; i <= Q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func add_interval(q, L, R int) {
	if L == R {
		return
	}
	if q != -1 {
		k := get_sum(L)
		a := find_kth_element(k)
		b := right_end[a]
		c := find_kth_element(k + 1)
		event[e][0] = q
		event[e][1] = b
		event[e][2] = L
		event[e][3] = c
		event[e][4] = -1
		e++
		event[e][0] = q
		event[e][1] = R
		event[e][2] = L
		event[e][3] = c
		event[e][4] = 1
		e++
	}
	is_left_end[L] = true
	Add(L, 1)
	right_end[L] = R
}

func get_sum(i int) int {
	s := 0
	for i > 0 && i < N {
		s += Bit[i]
		i -= i & -i
	}
	return s
}

func find_kth_element(k int) int {
	x, sx, dx := 0, 0, 1
	for 2*dx < N {
		dx *= 2
	}
	for dx > 0 {
		y := x + dx
		if y < N {
			sy := sx + Bit[y]
			if sy < k {
				x = y
				sx = sy
			}
		}
		dx /= 2
	}
	return x + 1
}

func rm_interval(q, L, R int) {
	if L == R {
		return
	}
	is_left_end[L] = false
	Add(L, -1)
	right_end[L] = 0
	if q != -1 {
		k := get_sum(L)
		a := find_kth_element(k)
		b := right_end[a]
		c := find_kth_element(k + 1)
		event[e][0] = q
		event[e][1] = b
		event[e][2] = L
		event[e][3] = c
		event[e][4] = 1
		e++
		event[e][0] = q
		event[e][1] = R
		event[e][2] = L
		event[e][3] = c
		event[e][4] = -1
		e++
	}
}

func Add(i, x int) {
	for i < N {
		Bit[i] += x
		i += i & -i
	}
}

func lowerBoundEvent(a [][5]int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i][1] >= x
	})
	return idx
}

func lowerBoundLRS(a [][3]int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i][0] >= x
	})
	return idx
}

func resize(a *[][5]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, [5]int{0, 0, 0, 0, 0})
		}
	}
}
