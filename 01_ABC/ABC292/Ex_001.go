package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 5e5 + 7

type pair struct {
	x, y int
}

var p, pre [maxn]int

type segment struct {
	mx  [maxn << 2]int
	tag [maxn << 2]int
}

func (seg *segment) update(k int) {
	seg.mx[k] = max(seg.mx[k<<1], seg.mx[k<<1|1])
}

func (seg *segment) build(k, l, r int) {
	if l == r {
		seg.mx[k] = pre[l]
		return
	}
	mid := (l + r) >> 1
	seg.build(k<<1, l, mid)
	seg.build(k<<1|1, mid+1, r)
	seg.update(k)
}

func (seg *segment) pushdown(k int) {
	if seg.tag[k] == 0 {
		return
	}
	seg.tag[k<<1] += seg.tag[k]
	seg.mx[k<<1] += seg.tag[k]
	seg.tag[k<<1|1] += seg.tag[k]
	seg.mx[k<<1|1] += seg.tag[k]
	seg.tag[k] = 0
	return
}

func (seg *segment) add(k, l, r, L, R, addi int) {
	if l == L && r == R {
		seg.tag[k] += addi
		seg.mx[k] += addi
		return
	}
	mid := (l + r) >> 1
	seg.pushdown(k)
	if R <= mid {
		seg.add(k<<1, l, mid, L, R, addi)
	} else if L > mid {
		seg.add(k<<1|1, mid+1, r, L, R, addi)
	} else {
		seg.add(k<<1, l, mid, L, mid, addi)
		seg.add(k<<1|1, mid+1, r, mid+1, R, addi)
	}
	seg.update(k)
}

func (seg *segment) query(k, l, r int) pair {
	if l == r {
		return pair{seg.mx[k], l}
	}
	seg.pushdown(k)
	mid := (l + r) >> 1
	if seg.mx[k<<1] >= 0 {
		return seg.query(k<<1, l, mid)
	} else {
		return seg.query(k<<1|1, mid+1, r)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var st segment
	var sum int

	var n, b, q int
	fmt.Fscan(in, &n, &b, &q)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		pre[i] = pre[i-1] + p[i] - b
		sum += p[i]
	}
	st.build(1, 1, n)
	for i := 1; i <= q; i++ {
		var c, x int
		fmt.Fscan(in, &c, &x)
		sum += x - p[c]
		st.add(1, 1, n, c, n, x-p[c])
		p[c] = x
		if st.mx[1] >= 0 {
			pii := st.query(1, 1, n)
			ans := float64(pii.x) / float64(pii.y)
			ans += float64(b)
			fmt.Fprintf(out, "%.10f\n", ans)
		} else {
			ans := float64(sum) / float64(n)
			fmt.Fprintf(out, "%.10f\n", ans)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
