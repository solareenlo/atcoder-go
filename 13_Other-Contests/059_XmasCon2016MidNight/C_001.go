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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	vs := make([]int, n)
	copy(vs, a)
	sort.Ints(vs)
	vs = unique(vs)
	lw, up := 0, len(vs)
	for up-lw > 1 {
		mid := (lw + up) / 2
		cur, cnt := 0, 0
		for i := 0; i < n; i++ {
			if a[i] >= vs[mid] {
				cur++
			} else {
				cur--
			}
			if cur == 0 {
				cnt++
			}
		}
		if cur > 0 || (cur == 0 && cnt%2 == 0) {
			lw = mid
		} else {
			up = mid
		}
	}
	for i := 0; i < n; i++ {
		if a[i] >= vs[lw] {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}

	white := make([][]bool, n)
	for i := range white {
		white[i] = make([]bool, n+1)
	}
	black := make([][]bool, n)
	for i := range black {
		black[i] = make([]bool, n+1)
	}

	for i := 0; i < n; i++ {
		cur, cnt := 0, 0
		for j := i; j < n; j++ {
			if a[j] != 0 {
				cur++
			} else {
				cur--
			}
			if cur == 0 {
				cnt++
			}
			if cur > 0 || (cur == 0 && cnt%2 == 0) {
				white[i][j+1] = true
			}
			if cur < 0 || (cur == 0 && cnt%2 == 0) {
				black[i][j+1] = true
			}
		}
	}

	l, r := 0, n
	for r-l > 1 {
		if r-l > 1 {
			f := -1
			for mid := l + 1; mid < r; mid++ {
				if !black[l][mid] && !black[mid][r] {
					f = mid
				}
			}
			fmt.Fprintln(out, f-l)
			out.Flush()
			var res string
			fmt.Fscan(in, &res)
			if res == "R" {
				r = f
			} else {
				l = f
			}
		}
		if r-l > 1 {
			var mid int
			fmt.Fscan(in, &mid)
			mid += l
			if white[l][mid] {
				fmt.Fprintln(out, "R")
				out.Flush()
				r = mid
			} else if white[mid][r] {
				fmt.Fprintln(out, "L")
				out.Flush()
				l = mid
			}
		}
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
