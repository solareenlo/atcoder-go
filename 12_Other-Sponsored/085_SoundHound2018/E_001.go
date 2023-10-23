package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAXN = 1 << 17

var N, Q int
var S []string

var Lcnt, Rcnt int

var LF, RF, G *Fenwick

func ok(mid int) bool {
	max_open := LF.lower_bound(mid)
	min_close := RF.lower_bound(Rcnt - mid + 1)
	return max_open < min_close
}

func solve() int {
	limit := min(Lcnt, Rcnt)
	if limit == 0 {
		return 0
	}
	lo := 0
	hi := limit + 1
	for hi-lo > 1 {
		mid := (lo + hi) / 2
		if ok(mid) {
			lo = mid
		} else {
			hi = mid
		}
	}
	if lo == 0 {
		return 0
	}
	half := lo
	max_open := LF.lower_bound(half)
	min_close := RF.lower_bound(Rcnt - half + 1)

	lo = max_open + 1
	hi = min_close + 1
	for hi-lo > 1 {
		mid := (lo + hi) / 2
		if RF.Sum(max_open+1, mid) == mid-(max_open+1) {
			lo = mid
		} else {
			hi = mid
		}
	}
	center := lo
	lo = 0
	hi = half
	hi = min(hi, int(RF.sum(center)))
	hi = min(hi, LF.Sum(center, MAXN))
	hi++
	for hi-lo > 1 {
		mid := (lo + hi) / 2
		lp := LF.lower_bound(half + mid)
		rp := RF.lower_bound(Rcnt - (half + mid) + 1)
		a := LF.sum(rp)
		b := RF.Sum(lp+1, MAXN)
		if int(a)+b-(lp-rp) > 0 {
			lo = mid
		} else {
			hi = mid
		}
	}

	lp := LF.lower_bound(half + lo)
	rp := RF.lower_bound(Rcnt - (half + lo) + 1)
	left := half
	right := half
	if lo != 0 {
		left = min(left, LF.Sum(0, rp))
		right = min(right, RF.Sum(lp+1, MAXN))
	}
	cnt := half + lo
	ret := 0
	if lo > 0 {
		ret += G.Sum(rp, lp+1)
		ret += LF.Sum(rp, lp+1) * (left - rp)
		ret += RF.Sum(rp, lp+1) * (rp - left)
	}

	ret += left * (left - 1) / 2
	ret += right * (right + 1) / 2
	ret -= 2 * cnt * right

	return -ret
}

func update(i, sgn int) {
	if S[i] == "(" {
		Lcnt += sgn
		LF.Add(i, sgn)
		G.Add(i, sgn*i)
	} else {
		Rcnt += sgn
		RF.Add(i, sgn)
		G.Add(i, -sgn*i)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	LF = NewFenwick(MAXN)
	RF = NewFenwick(MAXN)
	G = NewFenwick(MAXN)

	fmt.Fscan(in, &N, &Q)
	var tmp string
	fmt.Fscan(in, &tmp)
	S = strings.Split(tmp, "")

	for i := 0; i < N; i++ {
		update(i, 1)
	}

	for i := 0; i < Q; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		update(x, -1)
		if S[x] == "(" {
			S[x] = ")"
		} else {
			S[x] = "("
		}
		update(x, 1)
		ans := solve()
		fmt.Fprintln(out, ans)
	}
}

type Fenwick struct {
	n    int
	data []uint
}

func NewFenwick(n int) *Fenwick {
	fen := &Fenwick{
		n:    n,
		data: make([]uint, n),
	}
	for idx := range fen.data {
		fen.data[idx] = 0
	}
	return fen
}

func (fen *Fenwick) Add(pos, x int) {
	if !(0 <= pos && pos < fen.n) {
		panic("")
	}
	pos++
	for pos <= fen.n {
		fen.data[pos-1] += uint(x)
		pos += pos & -pos
	}
}

func (fen *Fenwick) Sum(l, r int) int {
	if !(0 <= l && l <= r && r <= fen.n) {
		panic("")
	}
	return int(fen.sum(r) - fen.sum(l))
}

func (fen *Fenwick) sum(r int) uint {
	s := uint(0)
	for r > 0 {
		s += fen.data[r-1]
		r -= r & -r
	}
	return s
}

func (fen *Fenwick) lower_bound(x int) int {
	if fen.n == 0 {
		return 0
	}
	i := 0
	s := 0
	for k := 1 << log2(fen.n); k > 0; k >>= 1 {
		if i+k <= fen.n && s+int(fen.data[i+k-1]) < x {
			i += k
			s += int(fen.data[i-1])
		}
	}
	return i
}

func (fen *Fenwick) upper_bound(x int) int {
	if fen.n == 0 {
		return 0
	}
	i := 0
	s := 0
	for k := 1 << log2(fen.n); k > 0; k >>= 1 {
		if i+k <= fen.n && !(x < s+int(fen.data[i+k-1])) {
			i += k
			s += int(fen.data[i-1])
		}
	}
	return i
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
