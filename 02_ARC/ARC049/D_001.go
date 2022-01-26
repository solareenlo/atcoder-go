package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	QL  int
	QR  int
	QV  int
	tag = [200_002]int{}
)

func sgt_modify(idx, l, r, d, oxo int) {
	if l >= QL && r <= QR {
		tag[idx] ^= QV
		return
	}
	oxo ^= tag[idx]
	flipped := 1
	if oxo&(1<<d) == 0 {
		flipped = 0
	}
	m := (l + r) >> 1
	if m >= QL {
		sgt_modify(idx*2+flipped, l, m, d+1, oxo)
	}
	if m+1 <= QR {
		tmp := 0
		if flipped == 0 {
			tmp = 1
		}
		sgt_modify(idx*2+tmp, m+1, r, d+1, oxo)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	for i := 0; i < q; i++ {
		var op, a, b int
		fmt.Fscan(in, &op, &a, &b)
		if op == 1 {
			a--
			u := 1
			oxo := 0
			for j := 0; j < n; j++ {
				oxo ^= tag[u]
				tmp1, tmp2 := 1, 1
				if oxo&(1<<j) == 0 {
					tmp1 = 0
				}
				if a&(1<<(n-1-j)) == 0 {
					tmp2 = 0
				}
				u = u*2 + (tmp1 ^ tmp2)
			}
			fmt.Fprintln(out, u-(1<<n)+1)
		} else {
			for i := 0; i < n; i++ {
				QL = max(1<<i, a)
				QR = min((2<<i)-1, b)
				QV = 1 << i
				if QL <= QR {
					sgt_modify(1, 1<<i, (2<<i)-1, 0, 0)
				}
			}
		}
	}
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
