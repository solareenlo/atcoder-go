package main

import (
	"bufio"
	"fmt"
	"os"
)

var f, ind [200200]int

func find(x int) int {
	if f[x] == x {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}

func join(u, v int) bool {
	a1 := find(u)
	a2 := find(v)
	if a1 != a2 {
		f[a1] = a2
	}
	ind[v] += 1
	return ind[v] > 2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		f[i] = i
	}
	if m == 0 || n == m || n-m != 1 {
		fmt.Println("No")
		return
	}
	bre := false
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if join(x, y) {
			bre = true
		}
	}
	if bre {
		fmt.Println("No")
		return
	}
	o := find(1)
	for i := 1; i <= n; i++ {
		if find(i) != o {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
