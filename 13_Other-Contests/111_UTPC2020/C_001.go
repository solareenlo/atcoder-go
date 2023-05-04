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
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const INF = int(1e18)

	var n, d int
	fmt.Fscan(in, &n, &d)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	ans := make([]pair, 0)
	vs := make([]int, n)
	for i := range vs {
		vs[i] = INF
	}
	sum, cnt := 0, 0
	for i := n - 1; i >= 0; i-- {
		if i+d < n && vs[i+d] != INF {
			if cnt%2 == 0 {
				sum += vs[i+d]
			} else {
				sum -= vs[i+d]
			}
			cnt--
		}
		if 2*d-1 <= i {
			cur := sum
			if cnt%2 == 0 {
				cur += a[i]
			} else {
				cur -= a[i]
			}
			x := b[i] + cur
			ans = append(ans, pair{i + 1 - d, x})
			vs[i] = x
			sum = x - sum
			cnt++
		}
		cur := sum
		if cnt%2 == 0 {
			cur += a[i]
		} else {
			cur -= a[i]
		}
		a[i] = cur
	}
	if n < 2*d-1 {
		for j := 0; j < 2; j++ {
			tar := b[n-d] + a[n-d]
			ok := true
			for i := n - d; i < d; i++ {
				if tar-a[i] != b[i] {
					ok = false
					break
				}
			}
			if ok {
				break
			}
			if j == 1 {
				fmt.Fprintln(out, "No")
				return
			}
			ans = append(ans, pair{n - d, 0})
			for i := n - d; i < n; i++ {
				a[i] = -a[i]
			}
		}
	}
	Len := min(n, 2*d-1)
	vs1 := make([]int, 0)
	for i := 0; i < Len-d; i++ {
		vs1 = append(vs1, b[i]-a[i])
	}
	for i := d - 1; i < Len; i++ {
		vs1 = append(vs1, b[i]+a[i])
	}
	tmp := sub(vs1)
	ans = append(ans, tmp...)

	fmt.Fprintln(out, "Yes")
	fmt.Fprintln(out, len(ans))
	for _, a := range ans {
		fmt.Fprintln(out, a.x+1, a.x+d, a.y)
	}
}

func sub(a []int) []pair {
	n := len(a)
	h := n / 2
	z := make([]int, h)
	sum := 0
	for i := 0; i < h; i++ {
		if i%2 == 0 {
			z[i] = -sum + a[i]
			sum += z[i]
		} else {
			z[i] = sum - a[i]
			sum -= z[i]
		}
	}
	w := make([]int, h)
	sum = a[n-1]
	for i := h - 1; i >= 0; i-- {
		w[i] = sum + a[h+i]
		sum = w[i] - sum
	}
	x := make([]int, h)
	y := make([]int, h)
	for i := 0; i < h; i++ {
		if (z[i] & 1) != (w[i] & 1) {
			fmt.Println("No")
			os.Exit(0)
		}
		x[i] = (w[i] - z[i]) / 2
		y[i] = (w[i] + z[i]) / 2
	}
	ans := make([]pair, 0)
	for i := 0; i < h; i++ {
		ans = append(ans, pair{i, x[i]})
	}
	ans = append(ans, pair{h, a[n-1]})
	for i := h - 1; i >= 0; i-- {
		ans = append(ans, pair{i, y[i]})
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
