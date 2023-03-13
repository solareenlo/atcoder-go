package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 1 << 60
const T = 1 << 18

var xx, aa [100010]int
var cc, yy, bb [100010]int
var all_x map[int]int
var seg [2 * T]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var NN int
	fmt.Fscan(in, &NN)
	for i := 0; i < NN; i++ {
		fmt.Fscan(in, &xx[i], &aa[i])
	}
	var QQ int
	fmt.Fscan(in, &QQ)
	for i := 0; i < QQ; i++ {
		fmt.Fscan(in, &cc[i], &yy[i], &bb[i])
	}

	all_x = make(map[int]int)
	for i := 0; i < NN; i++ {
		all_x[xx[i]] = 0
	}
	for i := 0; i < QQ; i++ {
		all_x[yy[i]] = 0
	}

	keys := make([]int, 0, len(all_x))
	for k := range all_x {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		x[N] = k
		all_x[k] = N
		N++
	}

	Init()

	for i := 0; i < NN; i++ {
		update_seg(all_x[xx[i]], aa[i])
		add(aa[i], 1)
	}

	for i := 0; i < QQ; i++ {
		if cc[i] == 1 {
			update_seg(all_x[yy[i]], bb[i])
			add(bb[i], 1)
		} else {
			update_seg(all_x[yy[i]], -1)
			add(bb[i], -1)
		}
		ans := Func()
		fmt.Fprintln(out, ans)
	}
}

func Init() {
	for i := 0; i < 2*T; i++ {
		seg[i] = -1
	}
}

func update_seg(x, c int) {
	x += T
	seg[x] = c
	for {
		x /= 2
		if x == 0 {
			return
		}
		if seg[2*x] == -1 {
			seg[x] = seg[2*x+1]
		} else if seg[2*x+1] == -1 {
			seg[x] = seg[2*x]
		} else if seg[2*x] == seg[2*x+1] {
			seg[x] = seg[2*x]
		} else {
			seg[x] = -2
		}
	}
}

var tree1, tree2 [T]int

func add(c, sgn int) {
	for i := c; i < T; i = ((i) | (i + 1)) {
		tree1[i] += sgn
	}
	for i := c; i < T; i = ((i) | (i + 1)) {
		tree2[i] += sgn * c
	}
}

func Func() int {
	ans := INF
	ans = min(ans, FUNC(rightmost(1)))
	ans = min(ans, FUNC(leftmost(1)))
	c := get_median()
	ans = min(ans, FUNC(c))
	return ans
}

func rightmost(id int) int {
	if seg[id] == -1 {
		return -1
	}
	if seg[id] != -2 {
		return seg[id]
	}
	if seg[2*id+1] != -1 {
		return rightmost(2*id + 1)
	}
	return rightmost(2 * id)
}

func leftmost(id int) int {
	if seg[id] == -1 {
		return -1
	}
	if seg[id] != -2 {
		return seg[id]
	}
	if seg[2*id] != -1 {
		return leftmost(2 * id)
	}
	return leftmost(2*id + 1)
}

func get_median() int {
	total := sum1(T)
	low := 0
	high := T
	for high-low > 1 {
		mid := (low + high) / 2
		tmp := sum1(mid)
		if tmp >= total-tmp {
			high = mid
		} else {
			low = mid
		}
	}
	return low
}

var N int
var x, a [200010]int

func FUNC(c int) int {
	ans := 0
	ans += sum1(c)*c - sum2(c)
	ans += (sum2(T) - sum2(c)) - (sum1(T)-sum1(c))*c
	L := leftmost2(c, 1, 0, T)
	if L == INF {
		L = 0
	} else {
		L = min(x[L], 0)
	}
	R := rightmost2(c, 1, 0, T)
	if R == -INF {
		R = 0
	} else {
		R = max(x[R], 0)
	}
	if R == 0 {
		return ans - L
	}
	if L == 0 {
		return ans + R
	}
	return ans + R - L + min(R, -L)
}

func sum1(pos int) int {
	ans := 0
	for i := pos; i > 0; i = ((i) & (i - 1)) {
		ans += tree1[i-1]
	}
	return ans
}

func sum2(pos int) int {
	ans := 0
	for i := pos; i > 0; i = ((i) & (i - 1)) {
		ans += tree2[i-1]
	}
	return ans
}

func leftmost2(c, id, low, high int) int {
	if seg[id] == -1 || seg[id] == c {
		return INF
	}
	if high-low == 1 {
		return low
	}
	mid := (low + high) / 2
	if seg[2*id] == -1 || seg[2*id] == c {
		return leftmost2(c, 2*id+1, mid, high)
	}
	return leftmost2(c, 2*id, low, mid)
}

func rightmost2(c, id, low, high int) int {
	if seg[id] == -1 || seg[id] == c {
		return -INF
	}
	if high-low == 1 {
		return low
	}
	mid := (low + high) / 2
	if seg[2*id+1] == -1 || seg[2*id+1] == c {
		return rightmost2(c, 2*id, low, mid)
	}
	return rightmost2(c, 2*id+1, mid, high)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
