package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var a [20]int

func f(p, q int) int {
	if p > n {
		return 0
	}
	if a[p] == q {
		return f(p+1, q)
	}
	if a[p] == 2 {
		return f(p+1, 4-q) + pow(3, n-p)
	}
	return f(p+1, q) + pow(3, n-p)*2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= 3; i++ {
		var t int
		fmt.Fscan(in, &t)
		for t > 0 {
			t--
			var t1 int
			fmt.Fscan(in, &t1)
			a[t1] = i
		}
	}
	t := min(f(1, 1), f(1, 3))
	if t <= m {
		fmt.Println(t)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
