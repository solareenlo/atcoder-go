package main

import (
	"bufio"
	"fmt"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	if n == 1 {
		add(0, 0)
		fmt.Fprintln(out, 0)
		out.Flush()
		return
	}
	mi := 0
	k := 1
	add(0, 1)
	for i := 1; i < n; i++ {
		x := add(i, -k)
		if x == -k {
			rem(i)
		} else {
			rem(mi)
			mi = i
			k = -k
		}
	}
	rem(mi)
	ls := -1
	k = 1
	for i := 0; i < n; i++ {
		if i == mi {
			continue
		}
		x := add(i, k*(n+n-i))
		if x == k {
			k = -k
			ls = i
		}
	}
	k = -k
	l := 0
	r := n + n - ls
	for r-l > 1 {
		mid := (l + r) / 2
		rem(ls)
		x := add(ls, k*mid)
		if x == k {
			r = mid
		} else {
			l = mid
		}
	}
	rem(ls)
	add(ls, k*r)
	k = -k
	add(mi, 0)
	l = 0
	r = n + 1
	for r-l > 1 {
		mid := (l + r) / 2
		rem(mi)
		x := add(mi, k*mid)
		if x == k {
			r = mid
		} else {
			l = mid
		}
	}
	rem(mi)
	add(mi, k*l)
	fmt.Fprintln(out, 0)
	out.Flush()
}

func add(i, x int) int {
	fmt.Fprintln(out, 1, i+1, x)
	out.Flush()
	var s string
	fmt.Scan(&s)
	if s[0] == '-' {
		os.Exit(0)
	} else if s[0] == 'L' {
		return -1
	} else if s[0] == 'B' {
		return 0
	}
	return 1
}

func rem(i int) int {
	fmt.Fprintln(out, 2, i+1)
	out.Flush()
	var s string
	fmt.Scan(&s)
	if s[0] == '-' {
		os.Exit(0)
	} else if s[0] == 'L' {
		return -1
	} else if s[0] == 'B' {
		return 0
	}
	return 1
}
