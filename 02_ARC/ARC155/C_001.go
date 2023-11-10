package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200005

var n int
var a0, a1 [N]int

func solve(a []int) {
	flag := false
	for i := 1; i <= n-2; i++ {
		if a[i]%2+a[i+1]%2+a[i+2]%2 == 2 {
			flag = true
			break
		}
	}
	if !flag {
		for i := 1; i <= n; i++ {
			p := i
			for p <= n && a[p]%2 == 0 {
				p++
			}
			if p-i >= 3 {
				sort.Ints(a[i:p])
			}
			i = p
		}
	} else {
		tot := 0
		c0, c1 := 0, 0
		for i := 1; i <= n; i++ {
			if a[i]%2 != 0 {
				c1++
				a1[c1] = a[i]
			} else {
				c0++
				a0[c0] = a[i]
			}
		}
		for i := 1; i <= c0; i++ {
			tot++
			a[tot] = a0[i]
		}
		for i := 1; i <= c1; i++ {
			tot++
			a[tot] = a1[i]
		}
		sort.Ints(a[c0+1 : n+1])
		if c0 >= 3 {
			sort.Ints(a[1 : c0+1])
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	solve(a)
	solve(b)
	for i := 1; i <= n; i++ {
		if a[i] != b[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
