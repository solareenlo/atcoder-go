package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	if m == 1 {
		for i := 0; i < n-1; i++ {
			if a[i][0] >= a[i+1][0] {
				fmt.Println(-1)
				return
			}
		}
		fmt.Println(0)
		return
	}
	sinf := 1 << 30
	var ncr func(int, int) int
	ncr = func(n, r int) int {
		p := 1
		for i := 1; i <= r; i++ {
			p *= (n - r + i)
			p /= i
			p = min(p, sinf)
		}
		return p
	}
	var less func([]int, []int, int, bool) bool
	less = func(a, b []int, mvs int, dfl bool) bool {
		rb := make([]int, len(b))
		copy(rb, b)
		if mvs <= 15 {
			for j := 0; j < mvs; j++ {
				for i := 1; i < m; i++ {
					rb[i] += rb[i-1]
					rb[i] = min(rb[i], sinf)
				}
			}
		} else {
			to := make([]int, m)
			for i := 0; i < m; i++ {
				if to[i] == sinf {
					break
				}
				for j := i; j < m; j++ {
					to[j] += b[i] * ncr(mvs-1+j-i, j-i)
					to[j] = min(to[j], sinf)
					if to[j] == sinf {
						break
					}
				}
			}
			for j := 1; j < m; j++ {
				to[j] = max(to[j], to[j-1])
			}
			rb = to
		}
		for i := 0; i < m; i++ {
			if a[i] != rb[i] {
				return a[i] < rb[i]
			}
		}
		return dfl
	}
	done, ans := 0, 0
	for i := 0; i < n-1; i++ {
		if a[i][0] > a[i+1][0] {
			fmt.Println(-1)
			return
		}
		if a[i][0] < a[i+1][0] {
			done = 0
			continue
		}
		p := a[i][0]
		q := a[i][1]
		x := a[i+1][0]
		y := a[i+1][1]
		if p*done+q < y {
			done = 0
		} else {
			lo := 0
			hi := 1 << 32
			ans := 0
			for lo <= hi {
				mid := (lo + hi) >> 1
				if x*mid+y <= p*done+q {
					ans = mid
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			}
			for {
				b := ans
				if done < ans {
					if !less(a[i], a[i+1], ans-done, false) {
						ans++
					}
				} else {
					if less(a[i+1], a[i], done-ans, true) {
						ans++
					}
				}
				if b == ans {
					break
				}
			}
			done = ans
		}
		ans += done
	}
	fmt.Println(ans)
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
