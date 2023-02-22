package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	check := func(n, k, x int) bool {
		Mx := 0
		if n%2 == 0 {
			n--
		}
		for n != 0 {
			t := 0
			if (n/3+1)%2 == 0 {
				t = 1
			}
			tmp := n/3 + 1 + t
			if x >= n {
				Mx += (n-tmp)/2 + 1
			} else if x >= tmp {
				Mx += (x-tmp)/2 + 1
			}
			n /= 3
			x >>= 1
		}
		return Mx >= k
	}

	solve := func() int {
		var n, k int
		fmt.Fscan(in, &n, &k)
		if k > (n+1)/2 {
			return -1
		}
		l, r, ans := 1, n, n
		for l <= r {
			mid := (l + r) >> 1
			if check(n, k, mid) {
				ans = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return ans
	}

	var T int
	fmt.Fscan(in, &T)
	for i := 0; i < T; i++ {
		fmt.Fprintln(out, solve())
	}
}
