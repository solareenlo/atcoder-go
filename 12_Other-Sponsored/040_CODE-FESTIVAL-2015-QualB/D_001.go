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

	ans := make([]int, 100001)
	S := make([]int, 100001)
	C := make([]int, 100001)
	comp := make([]int, 0)
	tree := make([]int, 100001)

	update := func(idx, x int) {
		for i := idx; i <= N; i += i & -i {
			tree[i] += int(x)
		}
	}

	query := func(idx int) int {
		ret := int(0)
		for i := idx; i > 0; i &= i - 1 {
			ret += tree[i]
		}
		return ret
	}

	for n := 1; n <= N; n++ {
		fmt.Fscan(in, &S[n], &C[n])
		comp = append(comp, S[n])
	}

	sort.Ints(comp)

	for n := 1; n <= N; n++ {
		k := sort.SearchInts(comp, S[n]) + 1
		for C[n] > 0 {
			lo := k - 1
			hi := N
			for lo+1 < hi {
				mid := (lo + hi) / 2
				if query(mid)-query(k-1) < comp[mid]-comp[k-1] {
					hi = mid
				} else {
					lo = mid
				}
			}
			diff := C[n]
			if hi != N {
				diff = min(C[n], comp[hi]-comp[k-1]-(query(hi)-query(k-1)))
			}
			if C[n] == diff {
				ans[n] = comp[hi-1] + query(hi) - query(hi-1) + C[n] - 1
			}
			update(hi, diff)
			C[n] -= diff
		}
	}

	for n := 1; n <= N; n++ {
		fmt.Fprintln(out, ans[n])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
