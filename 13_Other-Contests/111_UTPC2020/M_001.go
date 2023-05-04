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

	const INF = int(9e18)

	var N int
	fmt.Fscan(in, &N)

	query := make([][]int, N)
	for i := range query {
		query[i] = make([]int, 4)
	}
	for i := 0; i < N; i++ {
		var c string
		fmt.Fscan(in, &c, &query[i][1], &query[i][2])
		if c == "P" {
			query[i][0] = 0
		} else {
			query[i][0] = 1
		}
		query[i][3] = i
	}

	P := make([][]int, 0)
	Q := make([][]int, 0)
	for i := 0; i < N; i++ {
		if query[i][0] == 0 {
			P = append(P, []int{query[i][1], query[i][2], query[i][3]})
		} else {
			Q = append(Q, []int{query[i][1], query[i][2], query[i][3]})
		}
	}
	if len(P) == 0 || len(Q) == 0 {
		for i := 0; i < N; i++ {
			fmt.Fprintln(out, 0)
		}
		return
	}

	event_1 := calc_lower(P)

	for i := 0; i < len(P); i++ {
		P[i][0] *= -1
		P[i][1] *= -1
	}

	min_x := INF
	min_y := INF
	for i := 0; i < len(P); i++ {
		min_x = min(min_x, P[i][0])
		min_y = min(min_y, P[i][1])
	}
	for i := 0; i < len(P); i++ {
		P[i][0] -= min_x
		P[i][1] -= min_y
	}

	event_2 := calc_lower(P)
	for i := 0; i < len(event_2); i++ {
		event_2[i][2] *= -1
		event_2[i][3] *= -1
	}

	event_3 := calc_lower(Q)

	for i := 0; i < len(Q); i++ {
		Q[i][0] *= -1
		Q[i][1] *= -1
	}
	min_x = INF
	min_y = INF
	for i := 0; i < len(Q); i++ {
		min_x = min(min_x, Q[i][0])
		min_y = min(min_y, Q[i][1])
	}
	for i := 0; i < len(Q); i++ {
		Q[i][0] -= min_x
		Q[i][1] -= min_y
	}

	event_4 := calc_lower(Q)
	for i := 0; i < len(event_4); i++ {
		event_4[i][2] *= -1
		event_4[i][3] *= -1
	}

	event := make([][]int, 0)
	event = append(event, event_1...)
	event = append(event, event_2...)
	event = append(event, event_3...)
	event = append(event, event_4...)

	argsort := make([]int, len(event))
	for i := 0; i < len(event); i++ {
		argsort[i] = i
	}
	sort.Slice(argsort, func(i, j int) bool {
		return event[argsort[i]][0] < event[argsort[j]][0]
	})
	sorted_event := make([][]int, len(event))
	for i, idx := range argsort {
		sorted_event[i] = event[idx]
	}
	event = sorted_event

	vectors := make([][]int, len(event))
	for i := 0; i < len(event); i++ {
		vectors[i] = []int{0, 0}
		vectors[i][0] = event[i][2]
		vectors[i][1] = event[i][3]
	}
	sortSlice2(vectors)
	vectors = Unique(vectors)
	vectors = sort_by_atan2(vectors)

	type pair struct{ x, y int }

	to_id := make(map[pair]int)
	for i, v := range vectors {
		p := pair{v[0], v[1]}
		to_id[p] = i
	}

	for i := 0; i < len(event); i++ {
		event[i][2] = to_id[pair{event[i][2], event[i][3]}]
	}
	for i := 0; i < len(event); i++ {
		event[i] = event[i][:3]
	}

	ans := minkowski(len(query), event, vectors)

	initMod()

	var p_first, q_first int
	for i := range query {
		if query[i][0] == 0 {
			p_first = i
			break
		}
	}
	for i := range query {
		if query[i][0] == 1 {
			q_first = i
			break
		}
	}
	max_idx := max(p_first, q_first)
	for i := 0; i < max_idx; i++ {
		ans[i] = 0
	}

	for i := 0; i < N; i++ {
		fmt.Fprintln(out, (ans[i]%MOD)*int(mint(8).inv())%MOD)
	}
}

type tuple struct {
	x, y, z int
}

func minkowski(Q int, event, vectors [][]int) []int {
	N := len(vectors)
	idx := make([]int, Q+1)
	for i := 0; i <= Q; i++ {
		idx[i] = lowerBound(event, []int{i, 0, 0})
	}

	seg_raw := make([]tuple, N)
	seg := make([]tuple, 2*N)

	ans := make([]int, Q)
	for q := 0; q < Q; q++ {
		for e := idx[q]; e < idx[q+1]; e++ {
			t := event[e][1]
			i := event[e][2]
			if t == 0 {
				seg_raw[i].x += vectors[i][0]
				seg_raw[i].y += vectors[i][1]
			} else {
				seg_raw[i].x -= vectors[i][0]
				seg_raw[i].y -= vectors[i][1]
			}
			set_val(seg, i, seg_raw[i].x, seg_raw[i].y)
		}
		ans[q] = fold(seg, 0, N).z
	}
	return ans
}

func set_val(seg []tuple, i, x, y int) {
	N := len(seg) / 2
	i += N
	seg[i] = tuple{x, y, 0}
	for i > 1 {
		i >>= 1
		a := seg[i<<1].x
		b := seg[i<<1].y
		c := seg[i<<1].z
		d := seg[(i<<1)|1].x
		e := seg[(i<<1)|1].y
		f := seg[(i<<1)|1].z
		seg[i] = seg_f(a, b, c, d, e, f)
	}
}

func seg_f(a, b, c, d, e, f int) tuple {
	return tuple{a + d, b + e, c + f + a*e - b*d}
}

func fold(seg []tuple, l, r int) tuple {
	vl := tuple{0, 0, 0}
	vr := tuple{0, 0, 0}
	N := len(seg) / 2
	l += N
	r += N
	for l < r {
		if l&1 != 0 {
			vl = seg_f(vl.x, vl.y, vl.z, seg[l].x, seg[l].y, seg[l].z)
			l++
		}
		if r&1 != 0 {
			r--
			vr = seg_f(seg[r].x, seg[r].y, seg[r].z, vr.x, vr.y, vr.z)
		}
		l >>= 1
		r >>= 1
	}
	return seg_f(vl.x, vl.y, vl.z, vr.x, vr.y, vr.z)
}

func lowerBound(a [][]int, x []int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i][0] >= x[0]
	})
	return idx
}

func sort_by_atan2(XY [][]int) [][]int {
	XY1 := make([][]int, len(XY))
	XY2 := make([][]int, len(XY))
	for i := range XY {
		XY1[i] = make([]int, len(XY[i]))
		copy(XY1[i], XY[i])
		XY2[i] = make([]int, len(XY[i]))
		copy(XY2[i], XY[i])
	}

	N := len(XY)
	K := 0
	for {
		if K == 0 {
			K = 1
		} else {
			K = 2 * K
		}
		if K >= N {
			break
		}
		for p := 0; p < N; p += K + K {
			var first_A, last_A int
			if p < N {
				first_A = p
			} else {
				first_A = len(XY1)
			}
			if p+K < N {
				last_A = p + K
			} else {
				last_A = len(XY1)
			}
			A := XY1[first_A:last_A]
			var first_B, last_B int
			if p+K < N {
				first_B = p + K
			} else {
				first_B = len(XY1)
			}
			if p+K+K < N {
				last_B = p + K + K
			} else {
				last_B = len(XY1)
			}
			B := XY1[first_B:last_B]
			var first_C, last_C int
			if p < N {
				first_C = p
			} else {
				first_C = len(XY2)
			}
			if p+K+K < N {
				last_C = p + K + K
			} else {
				last_C = len(XY2)
			}
			C := XY2[first_C:last_C]
			a := len(A)
			b := len(B)
			i := 0
			j := 0
			for k := 0; k < a+b; k++ {
				if i == a || (j != b && is_sm(B[j], A[i])) {
					C[k] = B[j]
					j++
				} else {
					C[k] = A[i]
					i++
				}
			}
		}
		XY1, XY2 = XY2, XY1
	}
	return XY1
}

func is_sm(a, b []int) bool {
	ax := a[0]
	ay := a[1]
	bx := b[0]
	by := b[1]
	if ay < 0 {
		return by >= 0 || ax*by-ay*bx > 0
	}
	if ay == 0 {
		return ax >= 0 && (by > 0 || (by == 0 && bx < 0))
	}
	return by >= 0 && (ax*by-ay*bx) > 0
}

func calc_lower(XYT [][]int) [][]int {
	N := len(XYT)
	sort_key := make([]int, N)
	for i := 0; i < N; i++ {
		sort_key[i] = (XYT[i][0] << 32) | XYT[i][1]
	}
	argsort := make([]int, N)
	for i := 0; i < N; i++ {
		argsort[i] = i
	}
	sort.Slice(argsort, func(i, j int) bool {
		return sort_key[argsort[i]] < sort_key[argsort[j]]
	})

	RANK := make([]int, N)
	for i := 0; i < N; i++ {
		RANK[argsort[i]] = i
	}

	bit := make([]int, N+1)

	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = XYT[argsort[i]][0]
		Y[i] = XYT[argsort[i]][1]
	}

	event := make([][]int, 4*N)
	for i := range event {
		event[i] = make([]int, 4)
	}
	_n := 0

	var add_vector func(int, int, int)
	add_vector = func(t, i, j int) {
		dx := X[j] - X[i]
		dy := Y[j] - Y[i]
		event[_n][0] = t
		event[_n][1] = 0
		event[_n][2] = dx
		event[_n][3] = dy
		_n++
	}

	var delete_vector func(int, int, int)
	delete_vector = func(t, i, j int) {
		dx := X[j] - X[i]
		dy := Y[j] - Y[i]
		event[_n][0] = t
		event[_n][1] = 1
		event[_n][2] = dx
		event[_n][3] = dy
		_n++
	}

	var is_lower func(int, int, int) bool
	is_lower = func(l, m, r int) bool {
		x1 := X[m] - X[l]
		y1 := Y[m] - Y[l]
		x2 := X[r] - X[l]
		y2 := Y[r] - Y[l]
		return y1*x2 < y2*x1
	}

	var r, l int
	for i := 0; i < N; i++ {
		t := XYT[i][2]
		n := RANK[i]
		n_pts := get_sum(bit, N)
		if n_pts == 0 {
			add(bit, n, 1)
			n_pts += 1
			continue
		}
		k := get_sum(bit, n)
		if 1 <= k && k < n_pts {
			l := find_kth_element(bit, k)
			r := find_kth_element(bit, k+1)
			if !is_lower(l, n, r) {
				continue
			}
			delete_vector(t, l, r)
		}
		r = n
		for k >= 2 {
			m := find_kth_element(bit, k)
			l := find_kth_element(bit, k-1)
			if is_lower(l, m, r) {
				break
			}
			delete_vector(t, l, m)
			add(bit, m, -1)
			n_pts -= 1
			k -= 1
		}
		k = get_sum(bit, n)
		if k >= 1 {
			l := find_kth_element(bit, k)
			add_vector(t, l, n)
		}
		l = n
		for n_pts-k >= 2 {
			m := find_kth_element(bit, k+1)
			r := find_kth_element(bit, k+2)
			if is_lower(l, m, r) {
				break
			}
			delete_vector(t, m, r)
			add(bit, m, -1)
			n_pts -= 1
		}
		if n_pts-k >= 1 {
			r := find_kth_element(bit, k+1)
			add_vector(t, n, r)
		}
		add(bit, n, 1)
	}
	resize(&event, _n)
	return event
}

func get_sum(bit []int, i int) int {
	s := 0
	for i > 0 {
		s += bit[i]
		i -= i & (-i)
	}
	return s
}

func add(bit []int, i, x int) {
	i += 1
	for i < len(bit) {
		bit[i] += x
		i += i & (-i)
	}
}

func find_kth_element(bit []int, k int) int {
	N := len(bit)
	x := 0
	sx := 0
	dx := 1
	for 2*dx < N {
		dx *= 2
	}
	for dx > 0 {
		y := x + dx
		if y < N {
			sy := sx + bit[y]
			if sy < k {
				x = y
				sx = sy
			}
		}
		dx /= 2
	}
	return x
}

func sortSlice2(tmp [][]int) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i][0] == tmp[j][0] {
			return tmp[i][1] < tmp[j][1]
		}
		return tmp[i][0] < tmp[j][0]
	})
}

func resize(a *[][]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, []int{0, 0, 0, 0})
		}
	}
}

func Unique(a [][]int) [][]int {
	occurred := map[[2]int]bool{}
	result := [][]int{}
	for i := range a {
		if occurred[[2]int{a[i][0], a[i][1]}] != true {
			occurred[[2]int{a[i][0], a[i][1]}] = true
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

type mint int

func (m mint) pow(p int) mint {
	return powMod(m, p)
}

func (m mint) inv() mint {
	return invMod(m)
}

func (m mint) div(n mint) mint {
	return divMod(m, n)
}

const MOD = 998244353
const VMAX = 200005

var fact, invf [VMAX]mint

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := mint(1); i < VMAX; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
}

func powMod(a mint, n int) mint {
	res := mint(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a mint) mint {
	return powMod(a, MOD-2)
}

func divMod(a, b mint) mint {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a mint) mint {
	b, u, v := mint(MOD), mint(1), mint(0)
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
